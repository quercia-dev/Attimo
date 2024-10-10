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

	// Check if the table exists
	if !tx.Migrator().HasTable(categoryName) {
		tx.Rollback()
		return fmt.Errorf("table %s does not exist", categoryName)
	}

	// Get the table schema
	columns, err := tx.Table(categoryName).Migrator().ColumnTypes(&struct{}{})
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get column types for table %s: %w", categoryName, err)
	}

	// Debug: Print column information
	fmt.Printf("Columns for table %s:\n", categoryName)
	for _, col := range columns {
		fmt.Printf("  Name: %s, Type: %s\n", col.Name(), col.DatabaseTypeName())
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
		if !datatype.ValidateCheck(value) {
			tx.Rollback()
			return fmt.Errorf("invalid value for column %s: %v", columnName, value)
		}

		validatedData[columnName] = value
	}

	// Debug: Print column names and their respective values before inserting
	fmt.Printf("Inserting row into table %s with the following values:\n", categoryName)
	for colName, val := range validatedData {
		fmt.Printf("  Column: %s, Value: %v\n", colName, val)
	}

	// Create the row
	result := tx.Table(categoryName).Create(validatedData)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert row into %s: %w", categoryName, result.Error)
	}

	return tx.Commit().Error
}
