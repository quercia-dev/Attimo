package database

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

// getMetadata retrieves the metadata of the database.
// Returns a pointer to the last metadata row and an error.
func getMetadata(d *gorm.DB) (*Metadata, error) {
	var metadata Metadata
	if err := d.Last(&metadata).Error; err != nil {
		return nil, fmt.Errorf("error: Failed to get current metadata: %w", err)
	}
	return &metadata, nil
}

// getDatatype retrieves a datatype from the database by index.
// Returns a pointer to the row as a Datatype struct and an error.
func getDatatype(d *gorm.DB, index int) (*Datatype, error) {
	var datatype Datatype
	// Assuming index is the primary key of the Datatype table
	result := d.First(&datatype, index)

	if result.Error != nil {
		// Check if the error is because no record was found
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("error: No datatype found with index %d", index)
		}
		return nil, fmt.Errorf("error: Failed to get datatype with index %d: %w", index, result.Error)
	}

	return &datatype, nil
}

func toDBdatatype(goDatatype string) (string, error) {
	switch goDatatype {
	case IntType:
		return "INTEGER", nil
	case StringType:
		return "TEXT", nil
	case BoolType:
		return "TEXT", nil
	case TimeType:
		return "TEXT", nil
	default:
		return "", fmt.Errorf("error: unknown %s datatype", goDatatype)
	}
}

func isValidIdentifier(s string) bool {
	// Only allow alphanumeric characters and underscores
	match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", s)
	return match
}

func getTypeError(s string) error {
	return fmt.Errorf("%s is not a valid identifier", s)
}

func testValidIdentifier(types ...string) error {
	for _, t := range types {
		if !isValidIdentifier(t) {
			return getTypeError(t)
		}
	}
	return nil
}

// isGormModelColumn checks if the given column name is part of gorm.Model
func isGormModelColumn(columnName string) bool {
	return columnName == "id" || columnName == "created_at" || columnName == "updated_at" || columnName == "deleted_at"
}

// Helper function to get datatype by column name
func getDatatypeByName(tx *gorm.DB, columnName string) (*Datatype, error) {
	var datatype Datatype
	result := tx.Where("name = ?", columnName).First(&datatype)
	if result.Error != nil {
		return nil, result.Error
	}
	return &datatype, nil
}
