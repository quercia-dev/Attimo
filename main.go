package main

import (
	"context"
	"fmt"
	"path/filepath"

	ctrl "Attimo/control"
	data "Attimo/database"
	log "Attimo/logging"

	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

// App struct
type App struct {
	ctx     context.Context
	logger  *log.Logger
	data    *data.Database
	control *ctrl.Controller
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	dbFolder := filepath.Join(".", "db")
	dbPath := filepath.Join(dbFolder, "attimo.db")

	// Initialize logger
	logger, err := log.GetTestLogger()
	if err != nil {
		runtime.LogFatal(ctx, "Could not create logger: "+err.Error())
		return
	}
	a.logger = logger

	// TODO Initialize view
	// here

	// Initialize database
	database, err := data.SetupDatabase(dbPath, logger)
	if err != nil {
		logger.LogErr("Could not create database %v", err)
		runtime.LogFatal(ctx, "Could not create database: "+err.Error())
		return
	}
	a.data = database

	// Initialize controller
	controller, err := ctrl.New(database, logger)
	if err != nil {
		logger.LogErr("Could not create controller %v", err)
		runtime.LogFatal(ctx, "Could not create controller: "+err.Error())
		return
	}
	a.control = controller

	// TODO Initialize view with controller
	// here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:     "Your App",
		Width:     1024,
		Height:    768,
		OnStartup: app.startup,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}
