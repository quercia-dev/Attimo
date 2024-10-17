package database

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testDbPath = "test.db"
)

func TestAddRow(t *testing.T) {
	// Set up test database
	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, testDbPath)

	db, err := SetupDatabase(dbPath)
	if err != nil {
		t.Fatalf(dbSetupErrorString, err)
	}
	defer func() {
		db.Close()
		os.Remove(dbPath)
	}()

	currentTime := time.Now().Format(dateFormat)

	// Test cases
	tests := []struct {
		name         string
		categoryName string
		data         RowData
		wantErr      bool
	}{
		{
			name:         "Valid General row",
			categoryName: "General",
			data: RowData{
				"Opened":   currentTime,
				"Closed":   currentTime,
				"Note":     "Test note",
				"Project":  "Test Project",
				"Location": "Test Location",
				"File":     dbPath, // Using the test db path as a file that exists
			},
			wantErr: false,
		},
		{
			name:         "Valid Contact row",
			categoryName: "Contact",
			data: RowData{
				"Opened": currentTime,
				"Closed": currentTime,
				"Note":   "Test contact note",
				"Email":  "test@example.com",
				"Phone":  "1234567890",
				"File":   dbPath,
			},
			wantErr: false,
		},
		{
			name:         "Invalid email in Contact",
			categoryName: "Contact",
			data: RowData{
				"Opened": currentTime,
				"Closed": currentTime,
				"Note":   "Test contact note",
				"Email":  "not-an-email",
				"Phone":  "1234567890",
				"File":   dbPath,
			},
			wantErr: true,
		},
		{
			name:         "Missing required field in Financial",
			categoryName: "Financial",
			data: RowData{
				"Opened":   currentTime,
				"Closed":   currentTime,
				"Note":     "Test financial note",
				"Location": "Test Location",
				// Missing Cost_EUR
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.AddRow(tt.categoryName, tt.data)

			if tt.wantErr {
				if err == nil {
					t.Errorf("AddRow() error = nil, wantErr %v", tt.wantErr)
					return
				}
				return
			}
			if err != nil {
				t.Errorf("AddRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify the row was actually added
			query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE ", tt.categoryName)
			var conditions []string
			var values []interface{}
			for key, value := range tt.data {
				conditions = append(conditions, fmt.Sprintf("%s = ?", key))
				values = append(values, value)
			}
			query += strings.Join(conditions, " AND ")

			var count int64
			err = db.DB.Raw(query, values...).Count(&count).Error
			if err != nil {
				t.Errorf("Failed to verify row: %v", err)
			}
			if count != 1 {
				t.Errorf("Expected 1 row, got %d", count)
			}
		})
	}
}

func TestAddRowNonexistentCategory(t *testing.T) {
	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, testDbPath)

	db, err := SetupDatabase(dbPath)
	if err != nil {
		t.Fatalf(dbSetupErrorString, err)
	}
	defer func() {
		db.Close()
		os.Remove(dbPath)
	}()

	err = db.AddRow("NonexistentCategory", RowData{"Field": "value"})
	if err == nil {
		t.Error("Expected error when adding row to nonexistent category")
	}
}

func TestDeleteRow(t *testing.T) {
	// Set up test database
	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, testDbPath)

	db, err := SetupDatabase(dbPath)
	if err != nil {
		t.Fatalf(dbSetupErrorString, err)
	}
	defer func() {
		db.Close()
		os.Remove(dbPath)
	}()

	currentTime := time.Now().Format(dateFormat)

	// Helper function to add a test row
	addTestRow := func(categoryName string, data RowData) error {
		return db.AddRow(categoryName, data)
	}

	// Test cases
	tests := []struct {
		name            string
		categoryName    string
		setupData       RowData
		deleteCondition map[string]interface{}
		wantErr         bool
	}{
		{
			name:         "Delete existing General row",
			categoryName: "General",
			setupData: RowData{
				"Opened":   currentTime,
				"Closed":   currentTime,
				"Note":     "Test note",
				"Project":  "Test Project",
				"Location": "Test Location",
				"File":     dbPath,
			},
			deleteCondition: map[string]interface{}{"Project": "Test Project"},
			wantErr:         false,
		},
		{
			name:         "Delete non-existent Contact row",
			categoryName: "Contact",
			setupData: RowData{
				"Opened": currentTime,
				"Closed": currentTime,
				"Note":   "Test contact note",
				"Email":  "test@example.com",
				"Phone":  "1234567890",
				"File":   dbPath,
			},
			deleteCondition: map[string]interface{}{"Email": "nonexistent@example.com"},
			wantErr:         true,
		},
		{
			name:         "Delete Financial row with multiple conditions",
			categoryName: "Financial",
			setupData: RowData{
				"Opened":   currentTime,
				"Closed":   currentTime,
				"Note":     "Test financial note",
				"Location": "Test Location",
				"Cost_EUR": "100.50",
			},
			deleteCondition: map[string]interface{}{"Location": "Test Location", "Cost_EUR": "100.50"},
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup: Add a test row
			err := addTestRow(tt.categoryName, tt.setupData)
			assert.NoError(t, err, "Failed to add test row")

			// Perform deletion
			err = db.DeleteRow(tt.categoryName, tt.deleteCondition)

			if tt.wantErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}
		})
	}
}

func TestDeleteRowNonexistentCategory(t *testing.T) {
	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, testDbPath)

	db, err := SetupDatabase(dbPath)
	if err != nil {
		t.Fatalf(dbSetupErrorString, err)
	}
	defer func() {
		db.Close()
		os.Remove(dbPath)
	}()

	err = db.DeleteRow("NonexistentCategory", map[string]interface{}{"Field": "value"})
	assert.Error(t, err, "Expected error when deleting from nonexistent category")
}

func TestEditRow(t *testing.T) {
	// Set up test database
	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, "test.db")

	db, err := SetupDatabase(dbPath)
	if err != nil {
		t.Fatalf(dbSetupErrorString, err)
	}
	defer func() {
		db.Close()
		os.Remove(dbPath)
	}()

	currentTime := time.Now().Format(dateFormat)

	// Helper function to add a test row
	addTestRow := func(categoryName string, data RowData) error {
		return db.AddRow(categoryName, data)
	}

	// Test cases
	tests := []struct {
		name          string
		categoryName  string
		setupData     RowData
		editCondition map[string]interface{}
		editData      RowData
		wantErr       bool
	}{
		{
			name:         "Edit existing General row",
			categoryName: "General",
			setupData: RowData{
				"Opened":   currentTime,
				"Closed":   currentTime,
				"Note":     "Test note",
				"Project":  "Test Project",
				"Location": "Test Location",
				"File":     dbPath,
			},
			editCondition: map[string]interface{}{"Project": "Test Project"},
			editData: RowData{
				"Project": "Edited Project",
			},
			wantErr: false,
		},
		{
			name:         "Edit non-existent Contact row",
			categoryName: "Contact",
			setupData: RowData{
				"Opened": currentTime,
				"Closed": currentTime,
				"Note":   "Test contact note",
				"Email":  "test@example.com",
				"Phone":  "1234567890",
				"File":   dbPath,
			},
			editCondition: map[string]interface{}{"Email": "nonexistent@example.com"},
			editData: RowData{
				"Email": "one@two.com",
			},
			wantErr: true,
		},
		{
			name:         "Edit Financial row with multiple conditions",
			categoryName: "Financial",
			setupData: RowData{
				"Opened":   currentTime,
				"Closed":   currentTime,
				"Note":     "Test financial note",
				"Location": "Test Location",
				"Cost_EUR": "100.50",
			},
			editCondition: map[string]interface{}{"Location": "Test Location", "Cost_EUR": "100.50"},
			editData: RowData{
				"Cost_EUR": "200.00",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup: Add a test row
			err := addTestRow(tt.categoryName, tt.setupData)
			assert.NoError(t, err, "Failed to add test row")

			// Perform edit
			err = db.EditRow(tt.categoryName, tt.editCondition, tt.editData)

			if tt.wantErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}
		})
	}
}
