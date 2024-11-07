package control

import (
	data "Attimo/database"
	log "Attimo/logging"
	"fmt"
)

type Controller struct {
	logger *log.Logger
	data   *data.Database
}

func New(data *data.Database, logger *log.Logger) (*Controller, error) {
	if data == nil || logger == nil {
		return nil, fmt.Errorf("data or logger is nil")
	}

	return &Controller{
		logger: logger,
		data:   data,
	}, nil
}

func (c *Controller) GetCategories(logger *log.Logger) ([]string, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	return c.data.GetCategories()
}

func (c *Controller) GetCategoryColumns(logger *log.Logger, category string) ([]string, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	return c.data.GetCategoryColumns(category)
}

func (c *Controller) CreateRow(logger *log.Logger, category string, values data.RowData) error {
	if logger == nil {
		return fmt.Errorf(log.LoggerNilString)
	}

	return c.data.CreateRow(category, values)
}
