package database

import (
	"database/sql"
	"fmt"
	"strings"
)

// CreateRow inserts a new row into a category table
func (db *Database) CreateRow(categoryName string, data RowData) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Validate input data
	if err := db.validateInputData(tx, categoryName, data); err != nil {
		return fmt.Errorf("data validation failed: %w", err)
	}

	// Prepare column and value placeholders
	columns := make([]string, 0, len(data))
	placeholders := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))

	for col, val := range data {
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		values = append(values, val)
	}

	// Construct and execute the INSERT query
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		categoryName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	result, err := tx.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("failed to insert row: %w", err)
	}

	itemID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	// Add to pending if category has Opened field
	columns, err = db.GetCategoryColumns(categoryName)
	if err != nil {
		return fmt.Errorf("failed to get columns: %w", err)
	}

	for _, col := range columns {
		if col == "Opened" {
			if err := db.addToPending(tx, categoryName, int(itemID)); err != nil {
				return fmt.Errorf("failed to add to pending: %w", err)
			}
			break
		}
	}

	return tx.Commit()
}

func (db *Database) CloseItem(category string, itemID int, closeDate string) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Update the main record
	updateQuery := fmt.Sprintf(`
        UPDATE %s 
        SET Closed = ?, 
            updated_at = CURRENT_TIMESTAMP
        WHERE id = ? 
        AND deleted_at IS NULL
    `, category)

	result, err := tx.Exec(updateQuery, closeDate, itemID)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no item found with id %d in category %s", itemID, category)
	}

	// Remove from pending tracking
	if err := db.removeFromPending(tx, category, itemID); err != nil {
		return err
	}

	return tx.Commit()
}

// ReadRow retrieves a single row from a category table
func (db *Database) ReadRow(categoryName string, id int) (RowData, error) {
	// Query for column names
	columnQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ? AND deleted_at IS NULL LIMIT 1", categoryName)
	rows, err := db.DB.Query(columnQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query row: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	// Create a slice of interface{} to hold the values of each column
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	if err := rows.Scan(valuePtrs...); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	// Create the result map
	result := make(RowData)
	for i, col := range columns {
		result[col] = values[i]
	}

	return result, nil
}

func (db *Database) UpdateRow(categoryName string, id int, data RowData) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Validate input data
	if err := db.validateInputData(tx, categoryName, data); err != nil {
		return fmt.Errorf("data validation failed: %w", err)
	}

	// Prepare SET clause and values
	setClauses := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data)+1) // +1 for the ID

	for col, val := range data {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", col))
		values = append(values, val)
	}

	setClauses = append(setClauses, "updated_at = CURRENT_TIMESTAMP")
	values = append(values, id)

	// Construct the UPDATE query
	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = ? AND deleted_at IS NULL",
		categoryName,
		strings.Join(setClauses, ", "),
	)

	result, err := tx.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("failed to update row: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// DeleteRow soft deletes a row by setting its deleted_at timestamp
func (db *Database) DeleteRow(categoryName string, id int) error {
	query := fmt.Sprintf(
		"UPDATE %s SET deleted_at = CURRENT_TIMESTAMP WHERE id = ? AND deleted_at IS NULL",
		categoryName,
	)

	result, err := db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete row: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// ListRows retrieves multiple rows from a category table, with pagination
func (db *Database) ListRows(categoryName string, filters RowData, page, pageSize int) ([]RowData, int, error) {
	// Build WHERE clause from filters
	whereClause := "deleted_at IS NULL"
	values := make([]interface{}, 0)

	if len(filters) > 0 {
		conditions := make([]string, 0, len(filters))
		for col, val := range filters {
			conditions = append(conditions, fmt.Sprintf("%s = ?", col))
			values = append(values, val)
		}
		whereClause += " AND " + strings.Join(conditions, " AND ")
	}

	// Get total matching row count
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s", categoryName, whereClause)
	var total int
	err := db.DB.QueryRow(countQuery, values...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %w", err)
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Build main query with pagination
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE %s ORDER BY id DESC LIMIT ? OFFSET ?",
		categoryName,
		whereClause,
	)
	values = append(values, pageSize, offset)

	// Execute query
	rows, err := db.DB.Query(query, values...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query rows: %w", err)
	}
	defer rows.Close()

	// Get columns
	columns, err := rows.Columns()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get columns: %w", err)
	}

	// Prepare result slice
	var result []RowData

	// Iterate through rows
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}

		// Create row map
		rowData := make(RowData)
		for i, col := range columns {
			rowData[col] = values[i]
		}
		result = append(result, rowData)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating rows: %w", err)
	}

	return result, total, nil
}
