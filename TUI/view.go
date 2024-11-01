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
	model     tea.Model
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
	return LogsModel()
}

func (tui *TUI) Init(control *ctrl.Controller) error {
	if control == nil {
		err := fmt.Errorf("%v", nilControllerString)
		tui.logger.LogErr("%v", err)
		return err
	}

	tui.control = control

	var err error
	tui.model, err = MainModel(tui.logger)
	if err != nil {
		tui.logger.LogErr("Could not get Main model running")
		return err
	}
	p := tea.NewProgram(tui.model)

	newModel, err := p.Run()

	if err != nil {
		tui.logger.LogErr("tea program ran into an error %v", err)
		return err
	}

	tui.logger.LogInfo("Exited program p into %v", newModel)
	return nil
}
