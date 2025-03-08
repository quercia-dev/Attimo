package main

import (
	"fmt"
	"path/filepath"

	ctrl "Attimo/control"
	data "Attimo/database"
	log "Attimo/logging"
)

func main() {
	dbFolder := filepath.Join(".", "db")
	dbPath := filepath.Join(dbFolder, "attimo.db")

	// view.GetLogger()
	logger, err := log.GetTestLogger()
	if err != nil {
		fmt.Println("Could not create logger", err)
		return
	}

	var db *data.Database
	db, err = data.SetupDatabase(dbPath, logger)
	if err != nil {
		logger.LogErr("Could not create database %v", err)
		return
	}

	_, err = ctrl.New(db, logger)
	if err != nil {
		logger.LogErr("Could not create controller %v", err)
		return
	}
}
