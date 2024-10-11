package database

import (
	"fmt"
)

// RowData: map of column names to their values
type RowData map[string]interface{}

// AddRow adds a new row to the specified category table
func (d *Database) AddRow(categoryName string, data RowData) error {
	// Start a transaction
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get the table schema
	var dest interface{}
	columns, err := tx.Table(categoryName).Migrator().ColumnTypes(dest)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get column types for table %s: %w", categoryName, err)
	}

	// Create a map to store validated data
	validatedData := make(map[string]interface{})

	// Validate each column's data
	for _, column := range columns {
		columnName := column.Name()

		// Skip the gorm.Model columns
		if isGormModelColumn(columnName) {
			continue
		}

		value, exists := data[columnName]
		if !exists {
			tx.Rollback()
			return fmt.Errorf("missing value for column %s", columnName)
		}

		// Get the datatype for this column
		datatype, err := getDatatypeByName(tx, columnName)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to get datatype for column %s: %w", columnName, err)
		}

		// Validate the value
		if !datatype.ValidateValue(value) {
			tx.Rollback()
			return fmt.Errorf("invalid value for column %s: %v", columnName, value)
		}

		validatedData[columnName] = value
	}

	// Create the row
	result := tx.Table(categoryName).Create(validatedData)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert row into %s: %w", categoryName, result.Error)
	}

	return tx.Commit().Error
}
