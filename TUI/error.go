package tui

import (
	log "Attimo/logging"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// newMessage creates a new message model with the given message and replies.
// the method is decoupled from view.go for development
func newMessage(message string, replies []string, logger *log.Logger) (*selectionModel, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}
	if len(message) == 0 {
		message = "Error"
	}
	if len(replies) == 0 {
		replies = []string{"OK"}
	}
	return newSelectionModel(message, replies, logger)
}

// communicateError creates a new message model with the given message and replies
// and runs the model in a tea program
// the method does not return any values and handles errors internally
func communicateError(logger *log.Logger, message string) {
	if logger == nil {
		return
	}
	model, err := newMessage(message, nil, logger)
	if err != nil {
		logger.LogErr("Could not create error model")
		return
	}
	p := tea.NewProgram(model)
	_, err = p.Run()
	if err != nil {
		logger.LogErr("tea program ran into an error %v", err)
	}

}
