package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"Attimo/logging"

	_ "github.com/mattn/go-sqlite3"
)

// SetupDatabase initializes a Database struct and opens the database at the given path.
// If the DB does not exist, it will create a new database with the default schema.
// Returns a pointer to the database object and an error.
func SetupDatabase(path string, logger *logging.Logger) (*Database, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logger.LogWarn("Directory '%s' does not exist. Creating directory.", dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			logger.LogErr("Failed to create directory: %v", err)
			return nil, fmt.Errorf("failed to create directory: %v", err)
		}
	}

	fileExists := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fileExists = false
		logger.LogWarn("Database file '%s' does not exist. Creating new database.", path)
		// Touch the file to create it
		file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			logger.LogErr("Failed to create database file: %v", err)
			return nil, fmt.Errorf("failed to create database file: %v", err)
		}
		file.Close()
	}

	// Open the database
	sqlDB, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.LogErr("Failed to connect to database: %v", err)
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	database := &Database{
		Path:   path,
		DB:     sqlDB,
		logger: logger,
	}

	// Enable foreign key support
	_, err = sqlDB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		logger.LogErr("Failed to enable foreign key support: %v", err)
		return nil, fmt.Errorf("failed to enable foreign key support: %v", err)
	}

	if !fileExists {
		if err := database.createDefaultDB(); err != nil {
			logger.LogErr("Failed to create default DB: %v", err)
			return nil, fmt.Errorf("failed to create default DB: %v", err)
		}
		logger.LogInfo("New database created with default schema.")
	}

	logger.LogInfo("Database connection established correctly")
	return database, nil
}

func (db *Database) Close() {
	if db.DB != nil {
		db.DB.Close()
	}
}

// createDefaultDB sets up the initial database schema
func (db *Database) createDefaultDB() error {
	// Begin a transaction
	tx, err := db.DB.Begin()
	if err != nil {
		db.logger.LogErr("Failed to begin transaction: %v", err)
		return fmt.Errorf("failed to begin transaction: %v", err)
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
		db.logger.LogErr("Failed to create metadata table: %v", err)
		return fmt.Errorf("failed to create metadata table: %v", err)
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
		db.logger.LogErr("Failed to create datatypes table: %v", err)
		return fmt.Errorf("failed to create datatypes table: %v", err)
	}

	// Populate default data
	err = populateDefaultDB(tx, db.logger)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	db.logger.LogInfo("Default database setup completed successfully")
	return tx.Commit()
}

// populateDefaultDB fills the database with initial data
func populateDefaultDB(tx *sql.Tx, logger *logging.Logger) error {
	// Insert version
	_, err := tx.Exec("INSERT INTO metadata (version) VALUES (?)", currentVersion)
	if err != nil {
		logger.LogErr("Failed to insert version: %v", err)
		return fmt.Errorf("failed to insert version: %v", err)
	}

	// Prepare insert statement for datatypes
	stmt, err := tx.Prepare(`
        INSERT INTO datatypes 
        (name, variable_type, completion_value, completion_sort, value_check) 
        VALUES (?, ?, ?, ?, ?)
    `)
	if err != nil {
		logger.LogErr("Failed to prepare datatype insert: %v", err)
		return fmt.Errorf("failed to prepare datatype insert: %v", err)
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
			logger.LogErr("Failed to insert datatype %s: %v", dt.Name, err)
			return fmt.Errorf("failed to insert datatype %s: %v", dt.Name, err)
		}
	}

	// Create default category tables
	logger.LogInfo("Creating default category tables")
	err = createCategoryTables(tx, getDefaultCategories())
	if err != nil {
		logger.LogErr("Failed to create category tables: %v", err)
		return fmt.Errorf("failed to create category tables: %v", err)
	}
	logger.LogInfo("Successfully created category tables")

	return nil
}
