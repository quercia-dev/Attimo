package control

import (
	"Attimo/database"
	log "Attimo/logging"
	"database/sql"
	"fmt"
)

func New(data *database.Database, logger *log.Logger) (*Controller, error) {
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

func (c *Controller) GetCategoryColumns(logger *log.Logger, category string, condition *ColumnCondition) ([]string, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	// Get base columns from database
	columns, err := c.data.GetCategoryColumns(category)
	if err != nil {
		return nil, fmt.Errorf("failed to get base categories: %w", err)
	}

	// return all if no condition
	if condition == nil {
		return columns, nil
	}

	// create sets for faster lookups
	includeSet := make(map[string]bool)
	for _, col := range condition.IncludeColumn {
		includeSet[col] = true
	}

	excludeSet := make(map[string]bool)
	for _, col := range condition.ExcludeColumn {
		excludeSet[col] = true
	}

	var filteredColumns []string

	// start a transaction for getting datatypes
	tx, err := c.data.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	for _, column := range columns {
		include := true

		// check if column should be excluded
		if len(condition.ExcludeColumn) > 0 && excludeSet[column] {
			include = false
			continue
		}

		// check if column should be included
		if len(condition.IncludeColumn) > 0 {
			if !includeSet[column] {
				include = false
				continue
			}
		}

		// get datattpe information for column
		datatype, err := database.GetDatatypeByName(tx, column)
		if err != nil {
			logger.LogWarn("Failed to get datatype for column %s: %v", column, err)
			continue
		}

		// check fill behavior if specified
		if include && condition.FillBehavior != "" {
			if datatype.FillBehavior != condition.FillBehavior {
				include = false
				continue
			}
		}

		// Check data type if specified
		if include && condition.DataType != "" {
			if datatype.VariableType != condition.DataType {
				include = false
				continue
			}
		}

		if include {
			filteredColumns = append(filteredColumns, column)
		}
	}

	// Verify all required columns are present
	if len(condition.IncludeColumn) > 0 {
		foundColumns := make(map[string]bool)
		for _, col := range filteredColumns {
			foundColumns[col] = true
		}

		for _, requiredCol := range condition.IncludeColumn {
			if !foundColumns[requiredCol] {
				return nil, fmt.Errorf("required column %s not found in category %s", requiredCol, category)
			}
		}
	}

	return filteredColumns, nil
}

func (c *Controller) OpenItem(logger *log.Logger, request OpenItemRequest) OpenItemResponse {
	if logger == nil {
		return OpenItemResponse{Success: false, Error: fmt.Errorf(log.LoggerNilString)}
	}

	if request.Category == "" {
		return OpenItemResponse{Success: false, Error: fmt.Errorf("category is empty")}
	}
	if len(request.Values) == 0 {
		return OpenItemResponse{Success: false, Error: fmt.Errorf("values is empty")}
	}

	condition := &ColumnCondition{
		FillBehavior: "open",
	}

	columns, err := c.GetCategoryColumns(logger, request.Category, condition)
	if err != nil {
		return OpenItemResponse{Success: false, Error: fmt.Errorf("failed to get columns: %w", err)}
	}

	validColumns := make(map[string]bool)
	for _, col := range columns {
		validColumns[col] = true
	}

	rowData := make(database.RowData)
	for column, value := range request.Values {
		if !validColumns[column] {
			return OpenItemResponse{Success: false, Error: fmt.Errorf("invalid column: %s", column)}
		}
		rowData[column] = value
	}

	for column, value := range rowData {
		datatype, err := c.GetColumnDatatype(logger, request.Category, column)
		if err != nil {
			return OpenItemResponse{Success: false, Error: fmt.Errorf("failed to get datatype for column %s: %w", column, err)}
		}

		if !datatype.ValidateCheck(value, logger) {
			return OpenItemResponse{Success: false, Error: fmt.Errorf("invalid value for column %s: %v", column, value)}
		}
	}

	// create row
	err = c.CreateRow(logger, request.Category, rowData)
	if err != nil {
		return OpenItemResponse{Success: false, Error: fmt.Errorf("failed to create row: %w", err)}
	}

	return OpenItemResponse{Success: true, Error: nil}
}

func (c *Controller) CreateRow(logger *log.Logger, category string, values database.RowData) error {
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
		opts.Filters = database.RowData{}
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

	datatype, err := database.GetDatatypeByName(tx, column)
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

func (c *Controller) GetData(logger *log.Logger, category string) ([]string, []map[string]string, error) {
	cols := []string{"id", "name", "date", "column"}
	data := []map[string]string{
		{
			"id":     "12345",
			"name":   "Alice",
			"date":   "2024-12-16",
			"column": "value1",
		},
		{
			"id":     "12346",
			"name":   "Bob",
			"date":   "2024-12-17",
			"column": "value2",
		},
		{
			"id":     "12347",
			"name":   "Charlie",
			"date":   "2024-12-18",
			"column": "value3",
		},
	}

	return cols, data, nil
}
