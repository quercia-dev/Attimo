package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	database "Attimo/database"
)

func main() {

	// set up logging
	logDir := filepath.Join(".", "logs")
	if err := database.InitLogging(logDir); err != nil {
		log.Fatalf("failed to init logging: %v", err)
	}

	dbFolder := filepath.Join(".", "test")
	dbPath := filepath.Join(dbFolder, "central_storage.db")

	if _, err := os.Stat(dbFolder); os.IsNotExist(err) {
		if err := os.Mkdir(dbFolder, os.ModePerm); err != nil {
			fmt.Printf("Failed to create dir: %v\n", err)
			return
		}
	} else if err != nil {
		fmt.Printf("Failed to check directory: %v\n", err)
		return
	}

	// if file exists, delete it
	if _, err := os.Stat(dbPath); err == nil {
		os.Remove(dbPath)
	}

	db, err := database.SetupDatabase(dbPath)

	if err != nil {
		fmt.Printf("Error: could not create db object. %v\n", err)
		return
	}
	defer db.Close()

	// Test data insertion
	currentTime := time.Now().Format("02-01-2006")
	testData := []struct {
		category string
		data     database.RowData
	}{
		{
			category: "General",
			data: database.RowData{
				"Opened":   currentTime,
				"Closed":   currentTime,
				"Note":     "Test General note",
				"Project":  "Test Project",
				"Location": "Test Location",
				"File":     dbPath,
			},
		},
		{
			category: "General",
			data: database.RowData{
				"Opened":   currentTime,
				"Closed":   currentTime,
				"Note":     "General note",
				"Project":  "Project",
				"Location": "Location",
				"File":     dbPath,
			},
		},
		{
			category: "Contact",
			data: database.RowData{
				"Opened": currentTime,
				"Closed": currentTime,
				"Note":   "Test Contact note",
				"Email":  "test@example.com",
				"Phone":  "1234567890",
				"File":   dbPath,
			},
		},
	}

	for _, test := range testData {
		err := db.AddRow(test.category, test.data)
		if err != nil {
			fmt.Printf("Error adding row to %s: %v\n", test.category, err)
		} else {
			fmt.Printf("Successfully added row to %s\n", test.category)
		}
	}

	fmt.Println("Database setup and test data insertion complete.")
}
