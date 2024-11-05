package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	log "Attimo/logging"
)

// SetupDatabase initializes a Database struct and opens the database at the given path.
// If the DB does not exist, it will create a new database with the default schema.
// Returns a pointer to the database object and an error.
func SetupDatabase(path string) (*Database, error) {
	// Check if the directory exists
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist: %v", err)
	}

	fileExists := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fileExists = false
		log.LogWarn("Database file '%s' does not exist. Creating a new database.", path)
	}

	// Open the database
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, log.LogErr("failed to connect to database: %v", err)
	}

	database := &Database{
		Path: path,
		DB:   db,
	}

	// Enable foreign key support
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, log.LogErr("failed to enable foreign key support: %v", err)
	}

	if !fileExists {
		if err := database.createDefaultDB(); err != nil {
			return nil, log.LogErr("failed to create default DB: %v", err)
		}
		log.LogInfo("New database created with default schema.")
	}

	log.LogInfo("Database connection established correctly")
	return database, nil
}

func (db *Database) Close() {
	if db.DB != nil {
		db.DB.Close()
	}
}

// createDefaultDB sets up the initial database schema
func (d *Database) createDefaultDB() error {
	// Begin a transaction
	tx, err := d.DB.Begin()
	if err != nil {
		return log.LogErr("failed to begin transaction: %v", err)
	}

	// Create Metadata table
	_, err = tx.Exec(`
		CREATE TABLE metadata (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			version TEXT NOT NULL
		)
	`)
	if err != nil {
		tx.Rollback()
		return log.LogErr("failed to create metadata table: %v", err)
	}

	// Create Datatype table
	_, err = tx.Exec(`
		CREATE TABLE datatypes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			variable_type TEXT NOT NULL,
			completion_value TEXT NOT NULL,
			completion_sort TEXT NOT NULL,
			value_check TEXT NOT NULL
		)
	`)
	if err != nil {
		tx.Rollback()
		return log.LogErr("failed to create datatypes table: %v", err)
	}

	// Populate default data
	err = populateDefaultDB(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit()
}

// populateDefaultDB fills the database with initial data
func populateDefaultDB(tx *sql.Tx) error {
	// Insert version
	_, err := tx.Exec("INSERT INTO metadata (version) VALUES (?)", currentVersion)
	if err != nil {
		return log.LogErr("Failed to insert version: %v", err)
	}

	// Prepare insert statement for datatypes
	stmt, err := tx.Prepare(`
		INSERT INTO datatypes 
		(name, variable_type, completion_value, completion_sort, value_check) 
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return log.LogErr("Failed to prepare datatype insert: %v", err)
	}
	defer stmt.Close()

	// Insert default datatypes
	datatypes := getDefaultDatatypes()
	for _, dt := range datatypes {
		_, err = stmt.Exec(
			dt.Name,
			dt.VariableType,
			dt.CompletionValue,
			dt.CompletionSort,
			dt.ValueCheck,
		)
		if err != nil {
			return log.LogErr("Failed to insert datatype %s: %v", dt.Name, err)
		}
	}

	// Create default category tables
	return createCategoryTables(tx, getDefaultCategories())
}
