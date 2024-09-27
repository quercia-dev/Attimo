package database

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database struct holds the path to the database and the GORM database connection.
type Database struct {
	Path   string
	DB     *gorm.DB
	config *gorm.Config
}

// Metadata struct holds the metadata of the database.
type Metadata struct {
	gorm.Model
	Version string `gorm:"not null"`
}

// Category struct holds the category information.
type Category struct {
	gorm.Model
	Name    string          `gorm:"not null"`
	Columns json.RawMessage `gorm:"not null"`
}

const (
	// DefaultVersion is the default version of the database.
	currentVersion = "0.0.1"
)

// Datatype struct holds the datatype information.
type Datatype struct {
	gorm.Model
	Name string `gorm:"not null"`
	// VariableType is a string that determines the type of the field
	// if the VariableType field is empty, the value will be stored as a string
	VariableType string `gorm:"not null"`
	// CompletionValue is a string that determines how to complete the value of the field
	// if the CompletionValue field is 'no', the value will not be completed and CompletionSort will be ignored
	CompletionValue string `gorm:"not null"`
	// CompletionSort is a string that determines how to sort the completion values
	// if the CompletionSort field is 'no', the values will be displayed in the order they were provided by the CompletionValue field
	CompletionSort string `gorm:"not null"`
	// ValueCheck is a string that determines how to validate the value of the field
	// if the ValueCheck field is empty, the value will be validated with the default validation function, taken from the VariableType field
	// if the ValueCheck field is 'no', the value will not be validated
	ValueCheck string `gorm:"not null"`
}

// SetupDatabase initializes a Database struct and opens the database at the given path.
// If the DB does not exist, it will create a new database with the default schema.
// Returns a pointer to the database object and an error.
func SetupDatabase(path string) (*Database, error) {
	d := &Database{Path: path}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Warning: Database file '%s' does not exist. Creating empty file.\n", path)
	} else {
		fmt.Println("Database file exists already.")
	}

	d.config = &gorm.Config{}

	db, err := gorm.Open(sqlite.Open(path), d.config)
	if err != nil {
		return nil, fmt.Errorf("error: failed to connect to database: %w", err)
	}

	d.DB = db
	d.config.Logger = d.DB.Logger.LogMode(3)

	if err := d.createDefaultDB(); err != nil {
		return nil, fmt.Errorf("error: failed to create default DB: %w", err)
	}

	fmt.Println("Database connection established correctly")
	return d, nil
}

// Takes a pointer to a database and closes it
func (d *Database) Close() {
	sqlDB, _ := d.DB.DB()
	sqlDB.Close()
}

func (d *Database) createDefaultDB() error {
	if err := d.DB.AutoMigrate(&Metadata{}, &Category{}, &Datatype{}); err != nil {
		return fmt.Errorf("error: failed to migrate database: %w", err)
	}

	tx := d.DB.Begin()
	err := populateDB(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func populateDB(tx *gorm.DB) error {

	if err := tx.Create(&Metadata{Version: currentVersion}).Error; err != nil {
		return fmt.Errorf("error: Failed to insert version: %w", err)
	}
	timeType := "time.Time"
	// Populate the database with default datatypes
	datatypes := []Datatype{
		{Name: "Opened", VariableType: timeType, CompletionValue: "date", CompletionSort: "last", ValueCheck: ""},
		{Name: "Closed", VariableType: timeType, CompletionValue: "date", CompletionSort: "last", ValueCheck: ""},
		{Name: "Note", VariableType: "string", CompletionValue: "no", CompletionSort: "", ValueCheck: ""},
		{Name: "Project", VariableType: "string", CompletionValue: "unique", CompletionSort: "last", ValueCheck: ""},
		{Name: "Person", VariableType: "string", CompletionValue: "unique", CompletionSort: "frequency", ValueCheck: ""},
		{Name: "Location", VariableType: "string", CompletionValue: "unique", CompletionSort: "last", ValueCheck: ""},
		{Name: "URL", VariableType: "string", CompletionValue: "no", CompletionSort: "", ValueCheck: "URL"},
		{Name: "Cost (EUR)", VariableType: "integer", CompletionValue: "no", CompletionSort: "", ValueCheck: ""},
		{Name: "Deadline", VariableType: timeType, CompletionValue: "date", CompletionSort: "last", ValueCheck: ""},
		{Name: "Rating", VariableType: "integer", CompletionValue: "{1,2,3,4,5}", CompletionSort: "frequency", ValueCheck: "in{1,2,3,4,5}"},
		{Name: "Email", VariableType: "string", CompletionValue: "unique", CompletionSort: "", ValueCheck: "mail_ping"},
		{Name: "Phone", VariableType: "string", CompletionValue: "no", CompletionSort: "", ValueCheck: "phone"},
		{Name: "File", VariableType: "string", CompletionValue: "file", CompletionSort: "", ValueCheck: "file_exists"},
	}

	for _, datatype := range datatypes {
		if err := tx.Create(&datatype).Error; err != nil {
			return fmt.Errorf("error: Failed to insert datatype: %w", err)
		}
	}

	return nil
}
