package tui

import "github.com/charmbracelet/bubbles/textinput"

type Status int

const (
	StatusNone Status = iota
	StatusSuccess
	StatusError
)

type inputModel struct {
	tuiWindow
	prompt     string
	input      textinput.Model
	value      string
	status     Status
	statusMsg  string
	showStatus bool
}
