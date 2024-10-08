package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
}
