package control

import (
	log "Attimo/logging"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
)

func (c *Controller) setupServer(logger *log.Logger) error {
	if logger == nil {
		return fmt.Errorf(log.LoggerNilString)
	}

	app := fiber.New()

	c.addRoutes(app, logger)

	go func() {
		if err := app.Listen(":0"); err != nil {
			logger.LogErr("Error starting app: %v\n", err)
			return
		}

	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh

	logger.LogInfo("\nBackend Shutdown...")
	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		logger.LogWarn("Error during graceful shutdown: %v\n", err)
		return err
	} else {
		logger.LogWarn("Server shut down gracefully.")
	}
	return nil
}

func (c *Controller) addRoutes(app *fiber.App, logger *log.Logger) {
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello, World 👋!")
	})

	app.Get("/categories", func(ctx fiber.Ctx) error {
		list, err := c.GetCategories(logger)
		if err != nil {
			return err
		} else {
			return ctx.JSON(list)
		}
	})

}
