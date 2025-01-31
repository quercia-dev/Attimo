// schema.go
package database

import (
	"database/sql"
	"fmt"
	"strings"
)

// createCategoryTables creates tables for each category
func createCategoryTables(tx *sql.Tx, categories []CategoryTemplate) error {
	for _, cat := range categories {
		// Retrieve column definitions
		columnDefs, err := getColumnDefinitions(tx, cat.ColumnsID)
		if err != nil {
			return fmt.Errorf("failed to get column definitions for category %s: %v", cat.Name, err)
		}

		// Create table with standard CRUD columns and dynamic columns
		createTableSQL := fmt.Sprintf(`
            CREATE TABLE %s (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                deleted_at DATETIME,
                %s
            )
        `, cat.Name, strings.Join(columnDefs, ",\n"))

		_, err = tx.Exec(createTableSQL)
		if err != nil {
			return fmt.Errorf("failed to create table %s: %v", cat.Name, err)
		}
	}
	return nil
}

// getColumnDefinitions retrieves column definitions for a category
func getColumnDefinitions(tx *sql.Tx, columnIDs []int) ([]string, error) {
	var columnDefs []string

	for _, colID := range columnIDs {
		// Retrieve datatype information
		var datatype Datatype
		err := tx.QueryRow(`
            SELECT name, variable_type, completion_value, completion_sort, value_check, fill_behavior
            FROM datatypes 
            WHERE id = ?
        `, colID).Scan(
			&datatype.Name,
			&datatype.VariableType,
			&datatype.CompletionValue,
			&datatype.CompletionSort,
			&datatype.ValueCheck,
			&datatype.FillBehavior,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve datatype for ID %d: %v", colID, err)
		}

		// Convert Go type to SQLite type
		sqliteType, err := toSQLiteType(datatype.VariableType)
		if err != nil {
			return nil, err
		}

		columnDefs = append(columnDefs, fmt.Sprintf("%s %s", datatype.Name, sqliteType))
	}

	return columnDefs, nil
}

// toSQLiteType converts Go types to SQLite types
func toSQLiteType(goType string) (string, error) {
	switch goType {
	case IntType:
		return "INTEGER", nil
	case StringType:
		return "TEXT", nil
	case BoolType:
		return "INTEGER", nil // SQLite uses INTEGER for boolean
	case TimeType:
		return "DATETIME", nil
	case FloatType:
		return "REAL", nil
	case csvType:
		return "TEXT", nil // Store as comma-separated string
	default:
		return "", fmt.Errorf("unsupported type: %s", goType)
	}
}

func (data *Database) GetCategories() ([]string, error) {
	// Query for table names
	rows, err := data.DB.Query(`
		SELECT name
		FROM sqlite_master
		WHERE type = 'table'
		  AND name NOT LIKE 'sqlite_%'
		  AND name NOT LIKE 'datatypes'
		  AND name NOT LIKE 'metadata'
		ORDER BY name
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query table names: %w", err)
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("failed to scan table name: %w", err)
		}
		if name == "pending" {
			continue
		}
		categories = append(categories, name)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return categories, nil
}

func (data *Database) GetCategoryColumns(categoryName string) ([]string, error) {
	// Query for column names
	rows, err := data.DB.Query(`
		SELECT name
		FROM pragma_table_info(?)
		WHERE name != 'id'
		  AND name != 'created_at'
		  AND name != 'updated_at'
		  AND name != 'deleted_at'
	`, categoryName)
	if err != nil {
		return nil, fmt.Errorf("failed to query column names: %w", err)
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("failed to scan column name: %w", err)
		}
		columns = append(columns, name)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return columns, nil
}
