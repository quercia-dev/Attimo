package control

import (
	"Attimo/database"
	data "Attimo/database"
	log "Attimo/logging"
	"database/sql"
	"fmt"
)

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

func (c *Controller) ListRows(logger *log.Logger, opts ListRowsOptions) (*ListRowsResult, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	if opts.Category == "" {
		return nil, fmt.Errorf("category is empty")
	}

	// Set defaults if not provided
	if opts.Page < 1 {
		opts.Page = 1
	}
	if opts.PageSize < 1 {
		opts.PageSize = DefaultPageSize
	}
	if opts.Filters == nil {
		opts.Filters = data.RowData{}
	}

	// Get rows from database with pagination
	rows, total, err := c.data.ListRows(opts.Category, opts.Filters, opts.Page, opts.PageSize)
	if err != nil {
		logger.LogErr("Failed to list rows for category %s: %v", opts.Category, err)
		return nil, fmt.Errorf("failed to list rows: %w", err)
	}

	// Calculate total pages
	totalPages := (total + opts.PageSize - 1) / opts.PageSize

	result := &ListRowsResult{
		Rows:        rows,
		TotalRows:   total,
		CurrentPage: opts.Page,
		TotalPages:  totalPages,
		PageSize:    opts.PageSize,
	}

	logger.LogInfo("Listed %d rows (page %d of %d) for category %s",
		len(rows), opts.Page, totalPages, opts.Category)

	return result, nil
}

func (c *Controller) GetColumnDatatype(logger *log.Logger, category, column string) (*database.Datatype, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	tx, err := c.data.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	datatype, err := data.GetDatatypeByName(tx, column)
	if err != nil {
		return nil, fmt.Errorf("failed to get datatype: %w", err)
	}

	return datatype, nil
}

func (c *Controller) BeginTransaction() (*sql.Tx, error) {
	return c.data.DB.Begin()
}

func (c *Controller) CloseItem(logger *log.Logger, category string, itemID int, closeDate string) error {
	return c.data.CloseItem(category, itemID, closeDate)
}

func (c *Controller) GetPendingPointers(logger *log.Logger) ([]string, error) {
	return c.data.GetPendingPointers()
}
