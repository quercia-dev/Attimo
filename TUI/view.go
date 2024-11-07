package tui

import (
	ctrl "Attimo/control"
	log "Attimo/logging"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	nilControllerString = "pointer to controller is nil"
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
	mainShortcuts := map[string]int{
		openShortcut:   0,
		closeShortcut:  1,
		agendaShortcut: 2,
		editShortcut:   3,
		logShortcut:    4,
	}

	model, err := newBoxModel(tui.logger, mainItems, mainShortcuts)
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

		switch newModel.selected {
		case 0:
			value, err := tui.selectCategory()
			if err != nil {
				tui.logger.LogErr("Could not get category %v", err)
				return err
			}
			tui.logger.LogInfo("Selected category %v", value)
		default:
			tui.logger.LogWarn("Unexpected selection %v", newModel.selected)
		}

	} else {
		tui.logger.LogWarn("Unexpected model return type %v", ok)
	}

	return nil
}

func (tui *TUI) selectCategory() (interface{}, error) {
	categories, err := tui.control.GetCategories(tui.logger)
	if err != nil {
		tui.logger.LogErr("Could not get categories %v", err)
		return nil, err
	}
	model, err := newSelectionModel("Select category", categories, tui.logger)
	if err != nil {
		tui.logger.LogErr("Could not get selection model %v", err)
		return nil, err
	}
	p := tea.NewProgram(model)
	newModel, err := p.Run()
	if err != nil {
		tui.logger.LogErr("tea program ran into an error %v", err)
		return nil, err
	}
	tui.logger.LogInfo("Selection model running")
	if newModel, ok := newModel.(selectionModel); ok {
		tui.logger.LogInfo("Exited program by picking %v", newModel.selected)
		return newModel.selected, nil
	} else {
		tui.logger.LogWarn("Unexpected model return type %v", ok)
		return nil, nil
	}
}
