package main

import (
	model "Attimo/database"
	"fmt"
)

func main() {
	var path string = "./test/central_storage.db"

	var database model.Database = model.Database{}
	err := database.SetupDatabase(path)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		defer database.Close()
		fmt.Println("Database opened successfully")
	}
}
