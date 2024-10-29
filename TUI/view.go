package tui

import (
	ctrl "Attimo/control"
	log "Attimo/logging"
)

type TUI struct {
	logger    *log.Logger
	logsmodel *logsmodel
	control   *ctrl.Controller
}

func New(logger *log.Logger, logsmodel *logsmodel) (*TUI, error) {
	return &TUI{
		logger:    logger,
		logsmodel: logsmodel,
	}, nil
}

func GetLogger() (*log.Logger, *logsmodel, error) {
	return LogsModel()
}

func (tui *TUI) Init(control *ctrl.Controller) {
	tui.control = control
}