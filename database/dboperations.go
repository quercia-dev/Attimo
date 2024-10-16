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
	if err := d.validateTable(categoryName); err != nil {
		return err
	}

	model := &CategoryModel{}

	err := d.DB.Transaction(func(tx *gorm.DB) error {
		if err := d.populateModel(tx, categoryName, data, model); err != nil {
			return err
		}

		if err := d.insertRow(tx, categoryName, model); err != nil {
			return err
		}

		if err := d.updateFields(tx, categoryName, model); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (d *Database) DeleteRow(categoryName string, condition map[string]interface{}) error {
	if err := d.validateTable(categoryName); err != nil {
		return err
	}

	err := d.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Table(categoryName).Where(condition).Delete(&CategoryModel{})
		if result.Error != nil {
			return fmt.Errorf("failed to delete row: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no rows found matching the condition in table %s", categoryName)
		}

		return nil
	})

	return err
}

func (d *Database) EditRow(categoryName string, condition map[string]interface{}, data RowData) error {
	if err := d.validateTable(categoryName); err != nil {
		return err
	}

	err := d.DB.Transaction(func(tx *gorm.DB) error {
		// First, validate the data
		if err := d.validateGivenFields(tx, categoryName, data); err != nil {
			return err
		}

		// Check if the row exists before attempting updates
		var count int64
		if err := tx.Table(categoryName).Where(condition).Count(&count).Error; err != nil {
			return fmt.Errorf("failed to check row existence: %w", err)
		}
		if count == 0 {
			return fmt.Errorf("no rows found matching the condition in table %s", categoryName)
		}

		// Update each field
		for field, value := range data {
			if err := tx.Table(categoryName).Where(condition).Update(field, value).Error; err != nil {
				return fmt.Errorf("failed to update field %s: %w", field, err)
			}
		}

		return nil
	})

	return err
}

// validateField validates a single field value against its datatype
func (d *Database) validateField(tx *gorm.DB, columnName string, value interface{}) error {
	datatype, err := getDatatypeByName(tx, columnName)
	if err != nil {
		return fmt.Errorf("failed to get datatype for column %s: %w", columnName, err)
	}

	if !datatype.ValidateCheck(value) {
		return fmt.Errorf("invalid value for column %s: %v", columnName, value)
	}

	return nil
}

// validateFields validates the provided fields without requiring all fields to be present
func (d *Database) validateGivenFields(tx *gorm.DB, categoryName string, data RowData) error {
	columnTypes, err := tx.Migrator().ColumnTypes(categoryName)
	if err != nil {
		return fmt.Errorf("failed to get column types: %w", err)
	}

	// Create a map for O(1) column lookup
	columnMap := make(map[string]bool)
	for _, column := range columnTypes {
		columnMap[column.Name()] = true
	}

	// Validate each provided field
	for field, value := range data {
		if !columnMap[field] {
			return fmt.Errorf("invalid field name: %s", field)
		}

		if err := d.validateField(tx, field, value); err != nil {
			return err
		}
	}

	return nil
}

// validateTable checks if the specified table exists
func (d *Database) validateTable(categoryName string) error {
	if !d.DB.Migrator().HasTable(categoryName) {
		return fmt.Errorf("table %s does not exist", categoryName)
	}
	return nil
}

// SetField sets a field in the CategoryModel dynamically
func (m *CategoryModel) setField(name string, value interface{}) {
	if m.Fields == nil {
		m.Fields = make(map[string]interface{})
	}
	m.Fields[name] = value
}

// populateModel populates a model with data, validating all required fields
// Validates and populates model for AddRow operation
func (d *Database) populateModel(tx *gorm.DB, categoryName string, data RowData, model *CategoryModel) error {
	columnTypes, err := tx.Migrator().ColumnTypes(categoryName)
	if err != nil {
		return fmt.Errorf("failed to get column types: %w", err)
	}

	for _, column := range columnTypes {
		columnName := column.Name()

		if isGormModelColumn(columnName) {
			continue
		}

		value, exists := data[columnName]
		if !exists {
			return fmt.Errorf("missing required value for column %s", columnName)
		}

		// Reuse validateField
		if err := d.validateField(tx, columnName, value); err != nil {
			return err
		}

		model.setField(columnName, value)
	}

	return nil
}

// insertRow inserts the model into the specified table
func (d *Database) insertRow(tx *gorm.DB, categoryName string, model *CategoryModel) error {
	result := tx.Table(categoryName).Create(model)
	if result.Error != nil {
		return fmt.Errorf("failed to insert row: %w", result.Error)
	}
	return nil
}

// updateFields updates the Fields in the database
func (d *Database) updateFields(tx *gorm.DB, categoryName string, model *CategoryModel) error {
	for fieldName, fieldValue := range model.Fields {
		if err := tx.Table(categoryName).Where("id = ?", model.ID).Update(fieldName, fieldValue).Error; err != nil {
			return fmt.Errorf("failed to update field %s: %w", fieldName, err)
		}
	}
	return nil
}
