package main

import (
	model "Attimo/database"
	"fmt"
	"os"
)

func main() {
	var path string = "./test/central_storage.db"

	// if file exists, delete it
	if _, err := os.Stat(path); err == nil {
		os.Remove(path)
	}
	database, err := model.SetupDatabase(path)

	if err != nil {
		fmt.Println("Error: could not create db object.", err)
	} else {
		defer database.Close()
		fmt.Println("Database connection established successfully")
	}
}
