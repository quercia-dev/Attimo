package database

import (
	"fmt"
	"os"
	"path/filepath"

	log "Attimo/logging"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database struct holds the path to the database and the GORM database connection.
type Database struct {
	Path   string
	DB     *gorm.DB
	config *gorm.Config
	logger *log.Logger
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
	// VariableType is a string that determines the type of the field in go
	// the VariableType field cannot be empty
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
func SetupDatabase(path string, logger *log.Logger) (*Database, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	d := &Database{Path: path, logger: logger}

	// Check if the directory exists
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist: %v", err)
	}

	fileExists := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fileExists = false
		d.logger.LogWarn("Database file '%s' does not exist. Creating a new database.", path)
	} else {
		d.logger.LogInfo("Database file already exists.")
	}

	d.config = &gorm.Config{}

	db, err := gorm.Open(sqlite.Open(path), d.config)
	if err != nil {
		return nil, d.logger.LogErr("failed to connect to database: %v", err)
	}

	d.DB = db

	if !fileExists {
		if err := d.createDefaultDB(); err != nil {
			return nil, d.logger.LogErr("failed to create default DB: %v", err)
		}
		d.logger.LogInfo("New database created with default schema.")
	}

	d.logger.LogInfo("Database connection established correctly")
	return d, nil
}

// Takes a pointer to a database and closes it
func (d *Database) Close() {
	sqlDB, _ := d.DB.DB()
	sqlDB.Close()
}

func (d *Database) createDefaultDB() error {
	if err := d.DB.AutoMigrate(&Metadata{}, &Datatype{}); err != nil {
		return d.logger.LogErr("failed to migrate database: %v", err)
	}

	tx := d.DB.Begin()
	err := populateDefaultDB(tx, d.logger)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

// populateDB populates the database with the default schema.
func populateDefaultDB(tx *gorm.DB, logger *log.Logger) error {
	if logger == nil {
		return fmt.Errorf("nil logger")
	}

	if err := tx.Create(&Metadata{Version: currentVersion}).Error; err != nil {
		return logger.LogErr("Failed to insert version: %v", err)
	}

	datatypes := getDefaultDatatypes()
	if err := tx.Create(&datatypes).Error; err != nil {
		return logger.LogErr("Failed to insert datatypes: %v", err)
	}

	return addCategories(tx, getDefaultCategories(), logger)
}

func addCategories(tx *gorm.DB, categories []CategoryTemplate, logger *log.Logger) error {
	if logger == nil {
		return fmt.Errorf("nil logger")
	}

	for _, cat := range categories {
		// creates a new empty table inside the tx *gorm.DB with the structure of the Category struct
		if err := tx.Table(cat.Name).AutoMigrate(&Category{}); err != nil {
			return logger.LogErr(fmt.Sprintf("Failed to create table %s", cat.Name), err)
		}

		err := addColumns(tx, cat, logger)
		if err != nil {
			return err
		}
	}
	return nil
}

// addColumns adds columns to the category table, based on the columnsID field of the CategoryTemplate struct.
func addColumns(tx *gorm.DB, cat CategoryTemplate, logger *log.Logger) error {
	if logger == nil {
		return fmt.Errorf("nil logger")
	}

	for _, colID := range cat.ColumnsID {
		datatype, err := getDatatype(tx, colID)
		if err != nil {
			return logger.LogErr(fmt.Sprintf("Failed to retrieve datatype %d for category %s", colID, cat.Name), err)
		}

		if datatype.VariableType == "" {
			return logger.LogErr("Datatype %d has empty VariableType: %v", colID, fmt.Errorf("empty variable type"))
		}

		datatypeS, err := toDBdatatype(datatype.VariableType)
		if err != nil {
			return logger.LogErr("Failed to convert datatype for column %s in category %s: %v", datatype.Name, cat.Name, err)
		}

		err = testValidIdentifier(cat.Name, datatype.Name, datatypeS)
		if err != nil {
			return logger.LogErr("Failed to check identifier: %v", err)
		}

		command := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", cat.Name, datatype.Name, datatypeS)
		if err := tx.Exec(command).Error; err != nil {
			return logger.LogErr("Failed to add column %s to category %s: %v", datatype.Name, cat.Name, err)
		}
	}
	return nil
}
