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
	Path string
	DB   *gorm.DB
}

// Metadata struct holds the metadata of the database.
type Metadata struct {
	gorm.Model
	Version string
}

// Category struct holds the category information.
type Category struct {
	gorm.Model
	Name    string
	Columns json.RawMessage
}

// Datatype struct holds the datatype information.
type Datatype struct {
	gorm.Model
	Name            string
	VariableType    string
	CompletionValue string
	CompletionSort  string
	ValueCheck      string
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

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	d.DB = db

	if err := d.createDefaultDB(); err != nil {
		return nil, fmt.Errorf("Failed to create default DB: %w", err)
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
		return fmt.Errorf("Failed to migrate database: %w", err)
	}

	if err := d.populateDB(); err != nil {
		return fmt.Errorf("Failed to populate database: %w", err)
	}

	return nil
}

func (d *Database) populateDB() error {
	currentVersion := "0.0.1"

	if err := d.DB.Create(&Metadata{Version: currentVersion}).Error; err != nil {
		return fmt.Errorf("Failed to insert version: %w", err)
	}

	fmt.Println("Database populated successfully.")
	return nil
}

// runTransaction runs a transaction on the database.
// It takes a map of commands and runs them in a transaction.
// If any of the commands fail, it will rollback the transaction.
// Returns an error.
// func (d *Database) runTransaction(statements map[string]string) error {
// 	tx, err := d.db.Begin()
// 	if err != nil {
// 		return err
// 	}
//
// 	for key, statement := range statements {
// 		_, err := tx.Exec(statement)
// 		if err != nil {
// 			fmt.Printf("Error: Failed to execute command '%s': %v\nSQL: %s\n", key, err, statement)
// 			fmt.Println("Rolling Back transaction.")
// 			err := tx.Rollback()
// 			if err != nil {
// 				fmt.Println("Error: Failed to rollback transaction.", err)
// 			} else {
// 				fmt.Println("Transaction rolled back successfully.")
// 			}
// 			return err
// 		}
// 	}
// 	return tx.Commit()
// }
