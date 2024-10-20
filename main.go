package main

import (
	"fmt"
	"os"
	"path/filepath"

	log "Attimo/logging"

	TUI "Attimo/TUI"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	LogsModel := TUI.LogsModel()

	// set up logging
	if err := log.InitLoggingWithWriter(&LogsModel); err != nil {
		fmt.Println("Error: could not initialize logging.", err)
		return
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

	log.LogInfo("Starting TUI")

	p := tea.NewProgram(&LogsModel)
	if _, err := p.Run(); err != nil {
		log.LogErr("Error running TUI: %v", err)
	}

}
