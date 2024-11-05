package database

import (
	"Attimo/logging"
	"database/sql"
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
}

// Category struct holds the category information.
type Category struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Datatypes RowData
}

type CategoryTemplate struct {
	Name string
	// contains a list of numerical IDs for the rows of the datatypes
	ColumnsID []int
}

type RowData map[string]interface{}
