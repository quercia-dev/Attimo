package database

import (
	"fmt"

	"gorm.io/gorm"
)

// RowData: map of column names to their values
type RowData map[string]interface{}

// CategoryModel is a generic model for category tables
type CategoryModel struct {
	gorm.Model
	Fields map[string]interface{} `gorm:"-"` // Ignore this field in GORM
}

// AddRow adds a new row to the specified category table using GORM
func (d *Database) AddRow(categoryName string, data RowData) error {
	// Check if the table exists
	if !d.DB.Migrator().HasTable(categoryName) {
		return fmt.Errorf("table %s does not exist", categoryName)
	}

	// Create a new instance of the category model
	model := &CategoryModel{}

	// Transaction to insert
	err := d.DB.Transaction(func(tx *gorm.DB) error {

		columnTypes, err := tx.Migrator().ColumnTypes(categoryName)
		if err != nil {
			return fmt.Errorf("failed to get column types: %w", err)
		}

		for _, column := range columnTypes {
			columnName := column.Name()

			// Skip gorm.Model fields as they're handled automatically
			if isGormModelColumn(columnName) {
				continue
			}

			value, exists := data[columnName]
			if !exists {
				return fmt.Errorf("missing value for column %s", columnName)
			}

			// Get the datatype for this column
			datatype, err := getDatatypeByName(tx, columnName)
			if err != nil {
				return fmt.Errorf("failed to get datatype for column %s: %w", columnName, err)
			}

			// Validate the value
			if !datatype.ValidateCheck(value) {
				return fmt.Errorf("invalid value for column %s: %v", columnName, value)
			}

			model.setField(columnName, value)
		}
		// Create the new row
		result := tx.Table(categoryName).Create(model)
		if result.Error != nil {
			return fmt.Errorf("failed to insert row: %w", result.Error)
		}

		// If the insertion was successful, update the Fields in the database
		for fieldName, fieldValue := range model.Fields {
			if err := tx.Table(categoryName).Where("id = ?", model.ID).Update(fieldName, fieldValue).Error; err != nil {
				return fmt.Errorf("failed to update field %s: %w", fieldName, err)
			}
		}

		return nil
	})

	return err
}

// SetField sets a field in the CategoryModel dynamically
func (m *CategoryModel) setField(name string, value interface{}) {
	if m.Fields == nil {
		m.Fields = make(map[string]interface{})
	}
	m.Fields[name] = value
}
