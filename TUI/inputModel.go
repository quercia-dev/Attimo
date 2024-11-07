package tui

import (
	log "Attimo/logging"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type inputModel struct {
	tuiWindow
	prompt string
	input  textinput.Model
	value  string
}

func newInputModel(prompt string, logger *log.Logger) (*inputModel, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	ti := textinput.New()
	ti.Placeholder = alluringString
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return &inputModel{
		prompt:    prompt,
		input:     ti,
		tuiWindow: tuiWindow{logger: logger},
	}, nil
}

func (m inputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.HardQuit):
			m.logger.LogInfo("Quitting input")
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.Enter):
			m.value = m.input.Value()
			return m, tea.Quit
		}

		m.input, cmd = m.input.Update(msg)
		return m, cmd

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	return m, cmd
}

func (m inputModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.prompt,
		m.input.View(),
		"(enter to submit, ctrl+c to quit)",
	)
}
