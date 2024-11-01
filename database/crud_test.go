package database

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// TestDB is a helper struct for tests
type TestDB struct {
	*Database
	path string
}

// setupTestDB creates a temporary database for testing
func setupTestDB(t *testing.T) *TestDB {
	t.Helper()

	// Create temporary file for test database
	tmpfile, err := os.CreateTemp("", "test-db-*.sqlite")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Open database connection
	db := &Database{
		Path: tmpfile.Name(),
	}

	// Open the database connection
	sqlDB, err := sql.Open("sqlite3", db.Path)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	db.DB = sqlDB

	// Test the connection
	if err := db.DB.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	// Create required tables for testing
	if err := setupTestTables(db.DB); err != nil {
		t.Fatalf("Failed to setup test tables: %v", err)
	}

	return &TestDB{
		Database: db,
		path:     tmpfile.Name(),
	}
}

// setupTestTables creates the necessary tables for testing
func setupTestTables(db *sql.DB) error {
	// Create datatypes table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS datatypes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			variable_type TEXT NOT NULL,
			completion_value TEXT,
			completion_sort TEXT,
			value_check TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// Insert required datatypes
	_, err = db.Exec(`
		INSERT INTO datatypes (name, variable_type, value_check) VALUES 
		('Note', 'string', 'nonempty'),
		('Project', 'string', 'nonempty'),
		('Location', 'string', 'nonempty')
	`)
	if err != nil {
		return err
	}

	// Create General category table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS General (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			Note TEXT,
			Project TEXT,
			Location TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			deleted_at DATETIME DEFAULT NULL
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

// tearDown closes the database and removes the temporary file
func (tdb *TestDB) tearDown(t *testing.T) {
	t.Helper()

	if err := tdb.DB.Close(); err != nil {
		t.Errorf("Failed to close database: %v", err)
	}

	if err := os.Remove(tdb.path); err != nil {
		t.Errorf("Failed to remove test database file: %v", err)
	}
}

func TestCreateRow(t *testing.T) {
	db := setupTestDB(t)
	defer db.tearDown(t)

	tests := []struct {
		name         string
		categoryName string
		data         RowData
		wantErr      bool
	}{
		{
			name:         "valid row creation",
			categoryName: "General", // Opened, Closed, Note, Project, Location, File
			data: RowData{
				"Note":     "Test note",
				"Project":  "Test project",
				"Location": "Test location",
			},
			wantErr: false,
		},
		{
			name:         "invalid column",
			categoryName: "General", // Opened, Closed, Note, Project, Location, File
			data: RowData{
				"NonexistentColumn": "Test value",
			},
			wantErr: true,
		},
		{
			name:         "empty data",
			categoryName: "General",
			data:         RowData{},
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.CreateRow(tt.categoryName, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadRow(t *testing.T) {
	db := setupTestDB(t)
	defer db.tearDown(t)

	// Insert test data
	testData := RowData{
		"Note":     "Test note",
		"Project":  "Test project",
		"Location": "Test location",
	}

	err := db.CreateRow("General", testData)
	if err != nil {
		t.Fatalf("Failed to create test row: %v", err)
	}

	tests := []struct {
		name         string
		categoryName string
		id           int
		want         RowData
		wantErr      bool
	}{
		{
			name:         "existing row",
			categoryName: "General",
			id:           1,
			want:         testData,
			wantErr:      false,
		},
		{
			name:         "non-existent row",
			categoryName: "General",
			id:           999,
			want:         nil,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.ReadRow(tt.categoryName, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				for key, want := range tt.want {
					if got[key] != want {
						t.Errorf("ReadRow() got[%s] = %v, want %v", key, got[key], want)
					}
				}
			}
		})
	}
}

func TestUpdateRow(t *testing.T) {
	db := setupTestDB(t)
	defer db.tearDown(t)

	// Insert initial test data
	initialData := RowData{
		"Note":     "Initial note",
		"Project":  "Initial project",
		"Location": "Initial location",
	}

	err := db.CreateRow("General", initialData)
	if err != nil {
		t.Fatalf("Failed to create initial test row: %v", err)
	}

	tests := []struct {
		name         string
		categoryName string
		id           int
		data         RowData
		wantErr      bool
	}{
		{
			name:         "valid update",
			categoryName: "General",
			id:           1,
			data: RowData{
				"Note": "Updated note",
			},
			wantErr: false,
		},
		{
			name:         "non-existent row",
			categoryName: "General",
			id:           999,
			data: RowData{
				"Note": "Updated note",
			},
			wantErr: true,
		},
		{
			name:         "invalid column",
			categoryName: "General",
			id:           1,
			data: RowData{
				"NonexistentColumn": "Test value",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.UpdateRow(tt.categoryName, tt.id, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRow() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// Verify the update
				row, err := db.ReadRow(tt.categoryName, tt.id)
				if err != nil {
					t.Fatalf("Failed to read updated row: %v", err)
				}

				for key, want := range tt.data {
					if row[key] != want {
						t.Errorf("UpdateRow() updated value = %v, want %v", row[key], want)
					}
				}
			}
		})
	}
}

func TestDeleteRow(t *testing.T) {
	db := setupTestDB(t)
	defer db.tearDown(t)

	// Insert test data
	testData := RowData{
		"Note":     "Test note",
		"Project":  "Test project",
		"Location": "Test location",
	}

	err := db.CreateRow("General", testData)
	if err != nil {
		t.Fatalf("Failed to create test row: %v", err)
	}

	tests := []struct {
		name         string
		categoryName string
		id           int
		wantErr      bool
	}{
		{
			name:         "existing row",
			categoryName: "General",
			id:           1,
			wantErr:      false,
		},
		{
			name:         "non-existent row",
			categoryName: "General",
			id:           999,
			wantErr:      true,
		},
		{
			name:         "already deleted row",
			categoryName: "General",
			id:           1,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.DeleteRow(tt.categoryName, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRow() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// Verify the row is soft deleted
				row, err := db.ReadRow(tt.categoryName, tt.id)
				if err != sql.ErrNoRows {
					t.Errorf("DeleteRow() row still readable after deletion, got = %v", row)
				}
			}
		})
	}
}

func TestListRows(t *testing.T) {
	db := setupTestDB(t)
	defer db.tearDown(t)

	// Insert multiple test rows
	testData := []RowData{
		{"Note": "Note 1", "Project": "Project A", "Location": "Location X"},
		{"Note": "Note 2", "Project": "Project A", "Location": "Location Y"},
		{"Note": "Note 3", "Project": "Project B", "Location": "Location Z"},
	}

	for _, data := range testData {
		if err := db.CreateRow("General", data); err != nil {
			t.Fatalf("Failed to create test row: %v", err)
		}
	}

	tests := []struct {
		name         string
		categoryName string
		filters      RowData
		page         int
		pageSize     int
		wantCount    int
		wantErr      bool
	}{
		{
			name:         "list all rows",
			categoryName: "General",
			filters:      RowData{},
			page:         1,
			pageSize:     10,
			wantCount:    3,
			wantErr:      false,
		},
		{
			name:         "filter by project",
			categoryName: "General",
			filters:      RowData{"Project": "Project A"},
			page:         1,
			pageSize:     10,
			wantCount:    2,
			wantErr:      false,
		},
		{
			name:         "pagination",
			categoryName: "General",
			filters:      RowData{},
			page:         1,
			pageSize:     2,
			wantCount:    2,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows, total, err := db.ListRows(tt.categoryName, tt.filters, tt.page, tt.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(rows) != tt.wantCount {
					t.Errorf("ListRows() returned %d rows, want %d", len(rows), tt.wantCount)
				}

				if tt.filters["Project"] == "Project A" && total != 2 {
					t.Errorf("ListRows() total = %d, want 2 for Project A", total)
				}
			}
		})
	}
}
