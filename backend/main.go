package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	data "Attimo/database"
	log "Attimo/logging"

	"github.com/gofiber/fiber/v3"
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

	_, err = data.SetupDatabase(dbPath, logger)
	if err != nil {
		logger.LogErr("Could not create database %v", err)
		return
	}

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	go func() {
		if err := app.Listen(":0"); err != nil {
			fmt.Printf("Error starting app: %v\n", err)
		}

	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh

	fmt.Println("\nBackend Shutdown...")
	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		fmt.Printf("Error during graceful shutdown: %v\n", err)
	} else {
		fmt.Println("Server shut down gracefully.")
	}

}
