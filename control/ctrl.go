package control

import (
	data "Attimo/database"
	log "Attimo/logging"
)

type Controller struct {
	log  *log.Logger
	data *data.Database
}

func New(data *data.Database, logger *log.Logger) *Controller {
	return &Controller{
		log:  logger,
		data: data,
	}
}
