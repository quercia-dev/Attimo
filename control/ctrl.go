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
