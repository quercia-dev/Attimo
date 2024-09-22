package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Database struct holds the path to the database and the database connection.
type Database struct {
	path string
	db   *sql.DB
}

// exists checks if the database file exists.
// Returns a boolean.
func (d *Database) exists() bool {
	_, err := os.Stat(d.path)
	return !os.IsNotExist(err)
}

// SetupDatabase creates a new database object and opens the database at the given path.
// If the database does not exist, it will create a new database with the default schema.
// Returns a pointer to the database object and an error.
func SetupDatabase(path string) (*Database, error) {
	d := Database{}
	d.path = path
	err := d.open()

	if err != nil {
		return nil, err
	} else {
		if !d.exists() {
			fmt.Println("Warning: ", fmt.Sprintf("Database file '%s' does not exist.", d.path), "Creating empty file.")
			return &d, d.createDefaultDB()

		} else {
			fmt.Println("Database file exists already.")
			return &d, nil
		}
	}
}

// open runs the sql open command on the database path and returns the error.
func (d *Database) open() error {

	db, err := sql.Open("sqlite3", d.path)
	if err == nil {
		d.db = db
	}
	return err
}

// Close closes the database connection.
// Should be deferred after opening the database.
func (d *Database) Close() {
	d.db.Close()
}

// runTransaction runs a transaction on the database.
// It takes a map of commands and runs them in a transaction.
// If any of the commands fail, it will rollback the transaction.
// Returns an error.
func (d *Database) runTransaction(statements map[string]string) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	for key, statement := range statements {
		_, err := tx.Exec(statement)
		if err != nil {
			fmt.Printf("Error: Failed to execute command '%s': %v\nSQL: %s\n", key, err, statement)
			fmt.Println("Rolling Back transaction.")
			err := tx.Rollback()
			if err != nil {
				fmt.Println("Error: Failed to rollback transaction.", err)
			} else {
				fmt.Println("Transaction rolled back successfully.")
			}
			return err
		}
	}
	return tx.Commit()
}

// createEmptyDB creates an empty database with the default schema.
// Returns an error.
func (d *Database) createEmptyDB() error {
	statements := map[string]string{
		"metadataTable": `
		CREATE TABLE metadata (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        version TEXT NOT NULL);`,

		"categoriesTable": `
		CREATE TABLE categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		name TEXT NOT NULL, 
		columns TEXT NOT NULL
		)`,

		"datatypesTable": `
		CREATE TABLE datatypes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		variable_type TEXT NOT NULL,
		completion_value TEXT NOT NULL,
		completion_sort TEXT NOT NULL, 
		value_check TEXT NOT NULL DEFAULT 'DEFAULT'
		)`,
	}
	return d.runTransaction(statements)
}

// populateDB populates the database with the default values.
// for now, it only inserts the current version of the database.
// Returns an error.
func (d *Database) populateDB() error {
	currentVersion := "0.0.1"
	statements := map[string]string{
		"version": fmt.Sprintf(`INSERT INTO metadata (version) VALUES ('%s')`, currentVersion),

		"data types": `
		INSERT INTO datatypes (name, variable_type, completion_value, completion_sort, value_check) VALUES
		('Note', 'TEXT', 'none', 'frequency', 'DEFAULT'),
		('created', 'DATETIME', 'DATE', 'NONE', 'DEFAULT'),
		('updated', 'DATETIME', 'DATE', 'NONE', 'DEFAULT'),
		('opened', 'DATETIME', 'DATE', 'NONE', 'DEFAULT'),
		('closed', 'DATETIME', 'DATE', 'NONE', 'DEFAULT'),
		('File', 'TEXT', 'FILE', 'NONE', 'DEFAULT'),
		('Image', 'TEXT', 'FILE(IMAGE)', 'NONE', 'DEFAULT'),
		('Location', 'TEXT', 'OTHERS', 'FREQUENCY', 'DEFAULT'),
		('URL', 'TEXT', 'OTHERS', 'LENGTH', 'DEFAULT'),
		('Email', 'TEXT', 'EMAIL', 'FREQUENCY', 'DEFAULT'),
		('Phone', 'TEXT', 'PHONE', 'FREQUENCY', 'DEFAULT'),
		('Cost (USD)', 'TEXT', 'NUMBER', 'NONE', 'DEFAULT'),
		('Color', 'TEXT', 'COLOR', 'FREQUENCY', 'DEFAULT')`,

		// dummy data for now to test the database
		"categories": `
		INSERT INTO categories (name, columns) VALUES
		('General', 'created, updated, opened, closed, Note'),
		('Files', 'created, updated, opened, closed, File, Image'),
		('Contact', 'created, updated, opened, closed, Location, URL, Email, Phone'),
		('Finance', 'created, updated, opened, closed, Cost'),
		('Miscellaneous', 'created, updated, opened, closed, Color')`,
	}
	return d.runTransaction(statements)
}

// createDefaultDB creates an empty database with the default schema and populates it with the default values.
// it calls createEmptyDB and populateDB functions internally.
// Returns an error.
func (d *Database) createDefaultDB() error {
	err := d.createEmptyDB()
	if err != nil {
		fmt.Println("Error: Failed to create empty database.", err)
		return err
	} else {
		fmt.Println("Empty database created successfully.")
		err := d.populateDB()
		if err != nil {
			fmt.Println("Error: Failed to populate database.", err)
		} else {
			fmt.Println("Database populated successfully.")
		}
		return err
	}
}
