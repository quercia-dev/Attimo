package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// createPendingTable creates the pending tracking table
func createPendingTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
        CREATE TABLE IF NOT EXISTS pending (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            deleted_at DATETIME DEFAULT NULL,
            pointer TEXT NOT NULL UNIQUE,
            CHECK (pointer LIKE '%:%')
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create pending table: %w", err)
	}

	// Create index on pointer for faster lookups
	_, err = tx.Exec(`
        CREATE INDEX IF NOT EXISTS idx_pending_pointer 
        ON pending(pointer) 
        WHERE deleted_at IS NULL
    `)
	return err
}

// addToPending adds a new item to the pending tracking
func (db *Database) addToPending(tx *sql.Tx, category string, itemID int) error {
	pointer := fmt.Sprintf("%s:%d", category, itemID)

	_, err := tx.Exec(`
        INSERT INTO pending (pointer) 
        VALUES (?) 
        ON CONFLICT(pointer) DO UPDATE SET
        updated_at = CURRENT_TIMESTAMP,
        deleted_at = NULL
    `, pointer)

	if err != nil {
		return fmt.Errorf("failed to add to pending: %w", err)
	}
	return nil
}

// removeFromPending marks an item as no longer pending
func (db *Database) removeFromPending(tx *sql.Tx, category string, itemID int) error {
	pointer := fmt.Sprintf("%s:%d", category, itemID)

	_, err := tx.Exec(`
        UPDATE pending 
        SET deleted_at = CURRENT_TIMESTAMP,
            updated_at = CURRENT_TIMESTAMP
        WHERE pointer = ? 
        AND deleted_at IS NULL
    `, pointer)

	if err != nil {
		return fmt.Errorf("failed to remove from pending: %w", err)
	}
	return nil
}

// GetPendingItems returns all currently pending items with their details
func (db *Database) GetPendingItems() ([]RowData, error) {
	// Get all pending items
	rows, err := db.DB.Query(`
        SELECT id, pointer, created_at
        FROM pending
        WHERE deleted_at IS NULL
        ORDER BY created_at DESC
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to query pending items: %w", err)
	}
	defer rows.Close()

	var results []RowData
	for rows.Next() {
		var pendingID int
		var pointer string
		var createdAt time.Time

		if err := rows.Scan(&pendingID, &pointer, &createdAt); err != nil {
			return nil, fmt.Errorf("failed to scan pending row: %w", err)
		}

		// Parse pointer to get category and ID
		parts := strings.Split(pointer, ":")
		if len(parts) != 2 {
			db.logger.LogWarn("Invalid pointer format found: %s", pointer)
			continue
		}

		category, itemIDStr := parts[0], parts[1]

		// Get the item's details from its category table
		query := fmt.Sprintf(`
            SELECT %s.*
            FROM %s
            WHERE id = ? AND deleted_at IS NULL
        `, category, category)

		var itemData RowData
		itemData, err = db.queryRowToMap(query, itemIDStr)
		if err != nil {
			db.logger.LogWarn("Failed to get details for %s: %v", pointer, err)
			continue
		}

		// Add pending tracking information
		itemData["pending_id"] = pendingID
		itemData["category"] = category
		itemData["pending_since"] = createdAt

		results = append(results, itemData)
	}

	return results, nil
}

// Helper function to convert a row into a map
func (db *Database) queryRowToMap(query string, args ...interface{}) (RowData, error) {
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	for i := range values {
		valuePointers[i] = &values[i]
	}

	if err := rows.Scan(valuePointers...); err != nil {
		return nil, err
	}

	result := make(RowData)
	for i, col := range columns {
		result[col] = values[i]
	}

	return result, nil
}
