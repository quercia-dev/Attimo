package database

import (
	"fmt"
	"strings"
)

// RowData: map of column names to their values
type RowData map[string]interface{}

// AddRow adds a new row to the specified category table with SQL
func (d *Database) AddRow(categoryName string, data RowData) error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}

	// Start a transaction
	tx, err := sqlDB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Check if the table exists
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?", categoryName).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check if table exists: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("table %s does not exist", categoryName)
	}

	// Get the table schema
	rows, err := tx.Query(fmt.Sprintf("PRAGMA table_info(%s)", categoryName))
	if err != nil {
		return fmt.Errorf("failed to get table info: %w", err)
	}
	defer rows.Close()

	var columns []string
	var placeholders []string
	var values []interface{}

	for rows.Next() {
		var cid int
		var name, type_ string
		var notnull, pk int
		var dflt_value interface{}
		if err := rows.Scan(&cid, &name, &type_, &notnull, &dflt_value, &pk); err != nil {
			return fmt.Errorf("failed to scan column info: %w", err)
		}

		if isGormModelColumn(name) {
			continue
		}

		value, exists := data[name]
		if !exists {
			return fmt.Errorf("missing value for column %s", name)
		}

		// Get the datatype for this column
		datatype, err := getDatatypeByName(d.DB, name)
		if err != nil {
			return fmt.Errorf("failed to get datatype for column %s: %w", name, err)
		}

		// Validate the value
		if !datatype.ValidateCheck(value) {
			return fmt.Errorf("invalid value for column %s: %v", name, value)
		}

		columns = append(columns, name)
		placeholders = append(placeholders, "?")
		values = append(values, value)
	}

	// Prepare the INSERT statement
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		categoryName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	// Execute the INSERT statement
	_, err = tx.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("failed to insert row: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
