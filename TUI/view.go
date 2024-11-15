package tui

import (
	ctrl "Attimo/control"
	"Attimo/database"
	log "Attimo/logging"
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	nilControllerString = "pointer to controller is nil"
	valuePrompt         = "Enter value for %s:"
)

type TUI struct {
	logger    *log.Logger
	logsmodel *logsmodel
	control   *ctrl.Controller
}

func New(logger *log.Logger, logsmodel *logsmodel) (*TUI, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}
	return &TUI{
		logger:    logger,
		logsmodel: logsmodel,
	}, nil
}

func GetLogger() (*log.Logger, *logsmodel, error) {
	return newLogsModel()
}

func (tui *TUI) Init(control *ctrl.Controller) error {
	if control == nil {
		err := fmt.Errorf("%v", nilControllerString)
		tui.logger.LogErr("%v", err)
		return err
	}

	tui.control = control
	mainItems := []string{openItem, closeItem, agendaItem, editItem, logItem}

	model, err := newBoxModel(tui.logger, mainItems)
	if err != nil {
		tui.logger.LogErr("Could not get Main model")
		return err
	}
	p := tea.NewProgram(model)

	newModel, err := p.Run()
	if err != nil {
		tui.logger.LogErr("tea program ran into an error %v", err)
		return err
	}
	tui.logger.LogInfo("Main model running")

	if newModel, ok := newModel.(boxMenu); ok {
		tui.logger.LogInfo("Exited program by picking %v", newModel.selected)
		var err error
		switch newModel.selected {
		case 0: // OPEN
			err = tui.handleOpen()
		case 1: // CLOSE
			err = tui.handleClose()
		default:
			tui.logger.LogWarn("Unexpected selection %v", newModel.selected)
		}

		if err != nil {
			tui.logger.LogErr("Error handling open command: %v", err)
			return err
		} else {
			tui.logger.LogInfo("Successfully handled command")
		}

	} else {
		tui.logger.LogWarn("Unexpected model return type %v", ok)
	}

	return nil
}

func (tui *TUI) selectCategory() (string, error) {
	categories, err := tui.control.GetCategories(tui.logger)
	if err != nil {
		tui.logger.LogErr("Could not get categories %v", err)
		return "", err
	}

	model, err := newSelectionModel("Select category", categories, tui.logger)
	if err != nil {
		tui.logger.LogErr("Could not get selection model %v", err)
		return "", err
	}

	p := tea.NewProgram(model)
	newModel, err := p.Run()
	if err != nil {
		tui.logger.LogErr("tea program ran into an error %v", err)
		return "", err
	}

	tui.logger.LogInfo("Selection model running")

	if newModel, ok := newModel.(selectionModel); ok {
		if newModel.selected == nil {
			return "", fmt.Errorf("no category selected")
		}
		selectedIndex := newModel.selected.(int)
		if selectedIndex < 0 || selectedIndex >= len(categories) {
			return "", fmt.Errorf("invalid category index %v", selectedIndex)
		}
		selectedCategory := categories[selectedIndex]
		tui.logger.LogInfo("Selected category: %s", selectedCategory)
		return selectedCategory, nil
	}

	return "", fmt.Errorf("unexpected model return type")
}

func (tui *TUI) handleOpen() error {
	// Get and validate category selection
	category, err := tui.selectCategory()
	if err != nil {
		tui.logger.LogErr("Could not get category: %v", err)
		return err
	}

	// Get columns for selected category
	columns, err := tui.control.GetCategoryColumns(tui.logger, category)
	if err != nil {
		tui.logger.LogErr("Could not get columns: %v", err)
		return err
	}

	// asks the user to enter a row, and validates the data
	data, lastModel, err := tui.enterNewRowData(category, columns)
	if err != nil {
		tui.logger.LogErr("Could not get row data: %v", err)
	}

	// Create row and show final status
	if err := tui.control.CreateRow(tui.logger, category, data); err != nil {
		tui.logger.LogErr("Could not create row: %v", err)
		// Show error in the UI
		if lastModel != nil {
			p := tea.NewProgram(lastModel)
			lastModel.SetStatus(StatusError, fmt.Sprintf("Failed to create row: %v", err))
			p.Run()
		}
		return err
	}

	// Show success status
	if lastModel != nil {
		p := tea.NewProgram(lastModel)
		lastModel.SetStatus(StatusSuccess, "Row created successfully!")
		p.Run()
	}

	tui.logger.LogInfo("Row created successfully")
	return nil
}

func (tui *TUI) enterNewRowData(category string, columns []string) (database.RowData, *inputModel, error) {
	// create a map to store the data
	data := make(database.RowData)

	var lastModel *inputModel
	for _, column := range columns {
		if column == "Closed" {
			continue
		}

		// text input for each column
		model, err := newInputModel(fmt.Sprintf(valuePrompt, column), tui.logger)
		if err != nil {
			return nil, nil, fmt.Errorf("could not get input model for column %s: %w", column, err)
		}

		lastModel = model

		// loop until input is complete
		// input is complete when the user enters a valid or empty value
		var inputComplete bool
		for !inputComplete {
			p := tea.NewProgram(model)
			newModel, err := p.Run()

			if err != nil {
				return nil, nil, fmt.Errorf("tea program ran into an error for col %s: %w", column, err)
			}

			if inputModel, ok := newModel.(inputModel); ok {
				if inputModel.value == "" {
					model.SetStatus(StatusError, "Value insertion skipped")
					data[column] = ""
					inputComplete = true
				}

				// Validate the input value
				if err := tui.validateColumnInput(category, column, inputModel.value); err != nil {
					model.SetStatus(StatusError, fmt.Sprintf("Invalid input: %v", err))
					continue
				}

				data[column] = inputModel.value
				model.SetStatus(StatusSuccess, "Input accepted")
				inputComplete = true
			}
		}
	}
	return data, lastModel, nil
}

func (tui *TUI) validateColumnInput(category, column, value string) error {
	// Get the datatype for this column
	datatype, err := tui.control.GetColumnDatatype(tui.logger, category, column)
	if err != nil {
		return fmt.Errorf("failed to get datatype: %w", err)
	}

	// Create a transaction for validation
	tx, err := tui.control.BeginTransaction()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Validate the input
	if !datatype.ValidateCheck(value, tui.logger) {
		return fmt.Errorf("validation failed for %s", column)
	}

	return nil
}

func (tui *TUI) handleClose() error {
	// Create and run the close model
	model, err := newClosedModel(tui.logger, tui.control)
	if err != nil {
		if err.Error() == "no pending items to close" {
			// Show message to user that there are no pending items
			msgModel, err := newInputModel("No pending items to close. Press enter to continue.", tui.logger)
			if err != nil {
				return err
			}
			p := tea.NewProgram(msgModel)
			if _, err := p.Run(); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return fmt.Errorf("error running close model: %w", err)
	}

	if finalModel, ok := finalModel.(closedModel); ok {
		if finalModel.selected != "" && finalModel.timeInput.value != "" {
			// Split the pointer to get category and ID
			parts := strings.Split(finalModel.selected, ":")
			if len(parts) != 2 {
				return fmt.Errorf("invalid pointer format: %s", finalModel.selected)
			}

			category := parts[0]
			itemIDstring := parts[1]
			itemID, err := strconv.Atoi(itemIDstring)
			if err != nil {
				return fmt.Errorf("invalid item ID: %s", itemIDstring)
			}

			err = tui.control.CloseItem(tui.logger, category, itemID, finalModel.timeInput.value)
			if err != nil {
				tui.logger.LogErr("Failed to close item: %v", err)
				return fmt.Errorf("failed to close item: %w", err)
			}

			tui.logger.LogInfo("Successfully closed item %s in category %s", strconv.Itoa(itemID), category)
		}
	} else {
		return fmt.Errorf("unexpected model return type")
	}

	return nil
}
