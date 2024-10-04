package database

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

// RetrieveMetadata retrieves the metadata of the database.
// Returns a pointer to the last metadata row and an error.
func RetrieveMetadata(d *gorm.DB) (*Metadata, error) {
	var metadata Metadata
	if err := d.Last(&metadata).Error; err != nil {
		return nil, fmt.Errorf("error: Failed to get current metadata: %w", err)
	}
	return &metadata, nil
}

// RetrieveDatatype retrieves a datatype from the database by index.
// Returns a pointer to the row as a Datatype struct and an error.
func RetrieveDatatype(d *gorm.DB, index int) (*Datatype, error) {
	var datatype *Datatype
	if err := d.Find(&datatype, index).Error; err != nil {
		return nil, fmt.Errorf("error: Failed to get datatype: %w", err)
	}
	return datatype, nil
}

func DatabaseDatatype(goDatatype string) (string, error) {
	switch goDatatype {
	case "int":
		return "INTEGER", nil
	case "string":
		return "TEXT", nil
	case "bool":
		return "TEXT", nil
	case "time.Time":
		return "TEXT", nil
	default:
		return "", fmt.Errorf("error: unknown %s datatype", goDatatype)
	}
}

func IsValidIdentifier(s string) bool {
	// Only allow alphanumeric characters and underscores
	match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", s)
	return match
}

func ProduceTypeError(s string) error {
	return fmt.Errorf("%s is not a valid identifier", s)
}

func CheckValidIdentifier(types ...string) error {
	for _, t := range types {
		if !IsValidIdentifier(t) {
			return ProduceTypeError(t)
		}
	}
	return nil
}
