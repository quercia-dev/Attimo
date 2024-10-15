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

		if err := d.validateAndSetField(tx, columnName, data, model); err != nil {
			return err
		}
	}

	return nil
}

// validateAndSetField validates a single field and sets it in the model
func (d *Database) validateAndSetField(tx *gorm.DB, columnName string, data RowData, model *CategoryModel) error {
	value, exists := data[columnName]
	if !exists {
		return fmt.Errorf("missing value for column %s", columnName)
	}

	datatype, err := getDatatypeByName(tx, columnName)
	if err != nil {
		return fmt.Errorf("failed to get datatype for column %s: %w", columnName, err)
	}

	if !datatype.ValidateCheck(value) {
		return fmt.Errorf("invalid value for column %s: %v", columnName, value)
	}

	model.setField(columnName, value)
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
