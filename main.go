package main

import (
	"fmt"
	"os"
	"path/filepath"

	ctrl "Attimo/control"
	data "Attimo/database"
	view "Attimo/tui"
)

func main() {

	dbFolder := filepath.Join(".", "db")
	dbPath := filepath.Join(dbFolder, "attimo.db")

	// TEMPORARY: delete the database file so it resets every time
	// if file exists, delete it
	if _, err := os.Stat(dbPath); err == nil {
		os.Remove(dbPath)
	}

	logger, logsmodel, err := view.GetLogger()
	if err != nil {
		fmt.Println("Could not create logger", err)
		return
	}

	view, err := view.New(logger, logsmodel)
	if err == nil {
		fmt.Println("Could not create view", err)
	}

	data, err := data.SetupDatabase(dbPath, logger)
	if err != nil {
		fmt.Println("Could not create database", err)
		return
	}

	control := ctrl.New(data, logger)

	view.Init(control)
}
