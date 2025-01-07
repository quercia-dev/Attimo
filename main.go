package main

import (
	"fmt"
	"path/filepath"

	ctrl "Attimo/control"
	data "Attimo/database"
	log "Attimo/logging"
	view "Attimo/tui"
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

	view, err := view.New(logger, nil)
	if err != nil {
		logger.LogErr("Could not create view %v", err)
		return
	}

	data, err := data.SetupDatabase(dbPath, logger)
	if err != nil {
		logger.LogErr("Could not create database %v", err)
		return
	}

	control, err := ctrl.New(data, logger)
	if err != nil {
		logger.LogErr("Could not create controller %v", err)
		return
	}
	view.Init(control)
}
