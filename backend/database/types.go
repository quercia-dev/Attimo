package database

import (
	"Attimo/logging"
	"database/sql"
	"fmt"
	"time"
)

// Database struct holds the path to the database and the database connection.
type Database struct {
	Path   string
	DB     *sql.DB
	logger *logging.Logger
}

// Metadata struct holds the metadata of the database.
type Metadata struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

// Datatype struct holds the columns metadata.
type Datatype struct {
	ID              int
	Name            string
	VariableType    string
	CompletionValue string
	CompletionSort  string
	ValueCheck      string
	FillBehavior    string
}

// Entry struct holds the Entry information.
type Entry struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Datatypes RowData
}

type Timestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// Pending struct holds the pointer of an unclosed row.
type Pending struct {
	ID      int
	Pointer string // Format: "Category:ID", e.g., "General:123"
	Timestamp
}

// CategoryTemplate stores information to construct a new Category
type CategoryTemplate struct {
	Name string
	// contains a list of numerical IDs for the rows of the datatypes
	ColumnsID []int
}

// Category stores location and description of an existing Category
type Category struct {
	CategoryTemplate
	Timestamp
}

type RowData map[string]interface{}

func (row RowData) toString() map[string]string {
	result := make(map[string]string)

	for key, value := range row {
		strValue := fmt.Sprintf("%v", value)
		result[key] = strValue
	}
	return result
}

func RowDataToString(rows []RowData) ([]map[string]string, error) {
	result := make([]map[string]string, len(rows))

	for i, row := range rows {
		result[i] = row.toString()
	}
	return result, nil
}
