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

// Datatype struct holds the columns information.
type Datatype struct {
	ID              int
	Name            string
	VariableType    string
	CompletionValue string
	CompletionSort  string
	ValueCheck      string
	FillBehavior    string
}

// Category struct holds the category information.
type Category struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Datatypes RowData
}

// Pending struct holds the pointer of unclosed rows.
type Pending struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Pointer   string // Format: "Category:ID", e.g., "General:123"
}

type CategoryTemplate struct {
	Name string
	// contains a list of numerical IDs for the rows of the datatypes
	ColumnsID []int
}

type RowData map[string]interface{}

func (row RowData) toString() (map[string]string, error) {
	result := make(map[string]string)
	for key, value := range row {
		strValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("value of key %s is not a string", key)
		}
		result[key] = strValue
	}
	return result, nil
}

func RowDataToString(rows []RowData) ([]map[string]string, error) {
	result := make([]map[string]string, len(rows))
	for i, row := range rows {
		strRow, err := row.toString()
		if err != nil {
			return nil, fmt.Errorf("failed to convert row %d to string: %w", i, err)
		}
		result[i] = strRow
	}
	return result, nil
}
