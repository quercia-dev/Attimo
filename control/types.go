package control

import (
	data "Attimo/database"
	log "Attimo/logging"
)

const DefaultPageSize = 10

type Controller struct {
	logger *log.Logger
	data   *data.Database
}

type ListRowsOptions struct {
	Category string
	Filters  data.RowData // Optional
	Page     int          // Optional
	PageSize int          // Optional
}

type ListRowsResult struct {
	Rows        []data.RowData
	TotalRows   int
	CurrentPage int
	TotalPages  int
	PageSize    int
}
