package database

import (
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
}

// CategoryTemplate struct holds the category information.
// Not used in db but as a way to store the definition of category tables
type CategoryTemplate struct {
	Name string
	// contains a list of numerical GORM UID for the rows of the datatypes
	ColumnsID []int
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

	fileExists := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fileExists = false
		fmt.Printf("Warning: Database file '%s' does not exist. Creating a new database.\n", path)
	} else {
		fmt.Println("Database file already exists.")
	}

	d.config = &gorm.Config{}

	db, err := gorm.Open(sqlite.Open(path), d.config)
	if err != nil {
		return nil, fmt.Errorf("error: failed to connect to database: %w", err)
	}

	d.DB = db

	if !fileExists {
		if err := d.createDefaultDB(); err != nil {
			return nil, fmt.Errorf("error: failed to create default DB: %w", err)
		}
		fmt.Println("New database created with default schema.")
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
	if err := d.DB.AutoMigrate(&Metadata{}, &Datatype{}); err != nil {
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
		// 1
		{Name: "Opened", VariableType: timeType, CompletionValue: "date", CompletionSort: "last", ValueCheck: ""},
		// 2
		{Name: "Closed", VariableType: timeType, CompletionValue: "date", CompletionSort: "last", ValueCheck: ""},
		// 3
		{Name: "Note", VariableType: "string", CompletionValue: "no", CompletionSort: "", ValueCheck: ""},
		// 4
		{Name: "Project", VariableType: "string", CompletionValue: "unique", CompletionSort: "last", ValueCheck: ""},
		// 5
		{Name: "Person", VariableType: "string", CompletionValue: "unique", CompletionSort: "frequency", ValueCheck: ""},
		// 6
		{Name: "Location", VariableType: "string", CompletionValue: "unique", CompletionSort: "last", ValueCheck: ""},
		// 7
		{Name: "URL", VariableType: "string", CompletionValue: "no", CompletionSort: "", ValueCheck: "URL"},
		// 8
		{Name: "Cost (EUR)", VariableType: "integer", CompletionValue: "no", CompletionSort: "", ValueCheck: ""},
		// 9
		{Name: "Deadline", VariableType: timeType, CompletionValue: "date", CompletionSort: "last", ValueCheck: ""},
		// 10
		{Name: "Rating", VariableType: "integer", CompletionValue: "{1,2,3,4,5}", CompletionSort: "frequency", ValueCheck: "in{1,2,3,4,5}"},
		// 11
		{Name: "Email", VariableType: "string", CompletionValue: "unique", CompletionSort: "", ValueCheck: "mail_ping"},
		// 12
		{Name: "Phone", VariableType: "string", CompletionValue: "no", CompletionSort: "", ValueCheck: "phone"},
		// 13
		{Name: "File", VariableType: "string", CompletionValue: "file", CompletionSort: "", ValueCheck: "file_exists"},
	}

	if err := tx.Create(&datatypes).Error; err != nil {
		return fmt.Errorf("error: Failed to insert datatypes: %w", err)
	}

	categories := []CategoryTemplate{
		{Name: "General", ColumnsID: []int{1, 2, 3, 4, 6, 13}},
		{Name: "Contact", ColumnsID: []int{1, 2, 3, 11, 12, 13}},
		{Name: "Financial", ColumnsID: []int{1, 2, 3, 6, 8}},
	}

	for _, cat := range categories {
		// creates a new empty table inside the tx *gorm.DB with the structure of the Category struct
		err := tx.Table(cat.Name).AutoMigrate(&Category{})

		if err != nil {
			return fmt.Errorf("error: Failed to create table: %w", err)
		}

		// for each int in cat.ColumnsID, insert a column named after the ID int
		for _, colID := range cat.ColumnsID {
			columnName := fmt.Sprintf("Column%d", colID)
			err := tx.Exec("ALTER TABLE ? ADD COLUMN ? INTEGER", gorm.Expr(cat.Name), gorm.Expr(columnName)).Error
			if err != nil {
				return fmt.Errorf("error: Failed to insert column: %w", err)
			}
		}

	}

	return nil
}
