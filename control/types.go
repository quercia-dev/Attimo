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

type ColumnCondition struct {
	IncludeColumn []string
	ExcludeColumn []string
	FillBehavior  string // eg "open" or "close"
	DataType      string
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

type OpenItemRequest struct {
	Category string
	Values   map[string]string
}

type OpenItemResponse struct {
	Success bool
	Error   error
}

type ValidationResult struct {
	IsValid bool
	Message string
}

//type CloseItemResult struct {
//	Category  string
//	ItemID    int
//	CloseDate string
//	Error     error
//}
