package tui

import (
	ctrl "Attimo/control"
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
		case 3: // EDIT
			err = tui.handleEdit()
		default:
			tui.logger.LogWarn("Unexpected selection %v", newModel.selected)
		}

		if err != nil {
			tui.logger.LogErr("Error handling open command: %v", err)
			communicateError(tui.logger, fmt.Sprintf("Error handling open command: %v", err))
		} else {
			tui.logger.LogInfo("Successfully handled command")
			communicateError(tui.logger, "Successfully handled command")
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
	category, err := tui.selectCategory()
	if err != nil {
		tui.logger.LogErr("Could not get category: %v", err)
		return err
	}

	condition := &ctrl.ColumnCondition{
		FillBehavior: "open",
	}

	// Get columns for selected category
	columns, err := tui.control.GetCategoryColumns(tui.logger, category, condition)
	if err != nil {
		tui.logger.LogErr("Could not get columns: %v", err)
		return err
	}

	// collect values through user input
	values := make(map[string]string)
	for _, column := range columns {
		value, err := tui.promptForValue(column)
		if err != nil {
			tui.logger.LogErr("Could not get value for column %s: %v", column, err)
		}
		if value != "" { // only include non-empty values CHECK IF THIS IS NEEDED!!!
			values[column] = value
		}
	}

	// send request to controller
	request := ctrl.OpenItemRequest{
		Category: category,
		Values:   values,
	}

	response := tui.control.OpenItem(tui.logger, request)
	if response.Success {
		return response.Error
	}

	return nil
}

// helper function to get user input
func (tui *TUI) promptForValue(column string) (string, error) {
	model, err := newInputModel(fmt.Sprintf(valuePrompt, column), tui.logger)
	if err != nil {
		return "", fmt.Errorf("could not get input model for column %s: %w", column, err)
	}

	p := tea.NewProgram(model)
	result, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run input program: %w", err)
	}

	if m, ok := result.(inputModel); ok {
		return m.value, nil
	}

	return "", fmt.Errorf("unexpected model return type")
}

func (tui *TUI) handleClose() error {
	// Create and run the close model
	model, err := newClosedModel(tui.logger, tui.control)
	if err != nil {
		communicateError(tui.logger, fmt.Sprintf("Could not create close model: %v", err))
		return err
	}

	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return fmt.Errorf("error running close model: %w", err)
	}

	if finalModel, ok := finalModel.(*closedModel); ok {
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

func (tui *TUI) handleEdit() error {
	category, err := tui.selectCategory()
	if err != nil {
		tui.logger.LogErr("Could not get category: %v", err)
		return err
	}

	cols, rows, err := tui.control.GetData(tui.logger, category)
	if err != nil {
		tui.logger.LogErr("Could not get category data: %v", err)
		return err
	}

	model, err := newTableModel(tui.logger, cols, rows)
	if err != nil {
		tui.logger.LogErr("Could not get table model: %v", err)
		return err
	}

	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return fmt.Errorf("error running close model: %w", err)
	}

	if finalModel, ok := finalModel.(tableModel); ok {
		communicateError(tui.logger, fmt.Sprintf("Selected row: %v", finalModel.cursor))
		tui.logger.LogInfo("Selected row: %v", finalModel.cursor)
	}

	return nil
}
