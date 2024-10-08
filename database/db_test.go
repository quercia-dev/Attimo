package database

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupDatabase(t *testing.T) {
	// Create a temporary directory for test databases
	tempDir, err := os.MkdirTemp("", "test_db_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name    string
		dbPath  string
		wantErr bool
	}{
		{
			name:    "New database creation",
			dbPath:  filepath.Join(tempDir, "new.db"),
			wantErr: false,
		},
		{
			name:    "Invalid path",
			dbPath:  filepath.Join(tempDir, "invalid/path/db.db"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := SetupDatabase(tt.dbPath)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, db)
			assert.NotNil(t, db.DB)
			assert.Equal(t, tt.dbPath, db.Path)

			// Check if metadata was created correctly
			var metadata Metadata
			result := db.DB.First(&metadata)
			assert.NoError(t, result.Error)
			assert.Equal(t, currentVersion, metadata.Version)

			// Check if datatypes were created
			var datatypes []Datatype
			result = db.DB.Find(&datatypes)
			assert.NoError(t, result.Error)
			assert.NotEmpty(t, datatypes)

			db.Close()
		})
	}
}

func TestDatabase_Close(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "test_close.db")
	db, err := SetupDatabase(tempFile)
	assert.NoError(t, err)

	db.Close()

	// Verify that the connection is closed by attempting to ping
	sqlDB, err := db.DB.DB()
	assert.NoError(t, err)
	err = sqlDB.Ping()
	assert.Error(t, err)
}

func TestAddCategories(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "test_categories.db")
	db, err := SetupDatabase(tempFile)
	assert.NoError(t, err)
	defer db.Close()

	tests := []struct {
		name       string
		categories []CategoryTemplate
		wantErr    bool
	}{
		{
			name: "Valid categories",
			categories: []CategoryTemplate{
				{
					Name:      "test_category",
					ColumnsID: []int{1, 2}, // Assuming these IDs exist in datatypes
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid datatype ID",
			categories: []CategoryTemplate{
				{
					Name:      "invalid_category",
					ColumnsID: []int{9999}, // Non-existent ID
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := db.DB.Begin()
			err := addCategories(tx, tt.categories)
			if tt.wantErr {
				assert.Error(t, err)
				tx.Rollback()
				return
			}

			assert.NoError(t, err)
			tx.Commit()

			// Verify that the category table was created
			for _, cat := range tt.categories {
				assert.True(t, db.DB.Migrator().HasTable(cat.Name))
			}
		})
	}
}

func TestAddColumns(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "test_columns.db")
	db, err := SetupDatabase(tempFile)
	assert.NoError(t, err)
	defer db.Close()

	tests := []struct {
		name     string
		category CategoryTemplate
		wantErr  bool
	}{
		{
			name: "Valid columns",
			category: CategoryTemplate{
				Name:      "test_table",
				ColumnsID: []int{1}, // Assuming this ID exists in datatypes
			},
			wantErr: false,
		},
		{
			name: "Invalid datatype ID",
			category: CategoryTemplate{
				Name:      "invalid_table",
				ColumnsID: []int{9999}, // Non-existent ID
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := db.DB.Begin()
			// First create the table
			err := tx.Table(tt.category.Name).AutoMigrate(&Category{})
			assert.NoError(t, err)

			// Then add columns
			err = addColumns(tx, tt.category)
			if tt.wantErr {
				assert.Error(t, err)
				tx.Rollback()
				return
			}

			assert.NoError(t, err)
			tx.Commit()

			// Verify that the columns were added
			columns, err := db.DB.Migrator().ColumnTypes(tt.category.Name)
			assert.NoError(t, err)
			assert.GreaterOrEqual(t, len(columns), len(tt.category.ColumnsID))
		})
	}
}
