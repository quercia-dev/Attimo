package tui

import (
	log "Attimo/logging"
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Status int

const (
	StatusNone Status = iota
	StatusSuccess
	StatusError
)

type inputModel struct {
	tuiWindow
	keys       selectionKeyMap
	prompt     string
	input      textinput.Model
	value      string
	status     Status
	statusMsg  string
	showStatus bool
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
		tuiWindow: tuiWindow{
			help:   help.New(),
			logger: logger,
		},
		keys:       newSelectionKeyMap(),
		prompt:     prompt,
		input:      ti,
		status:     StatusNone,
		showStatus: false,
	}, nil
}

func (m inputModel) Init() tea.Cmd {
	return textinput.Blink
}

// SetStatus updates the status message of the input model
func (m *inputModel) SetStatus(status Status, msg string) {
	m.status = status
	m.statusMsg = msg
	m.showStatus = true
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.logger.LogInfo("Quitting input")
			return m, tea.Quit
		case key.Matches(msg, m.keys.Enter):
			m.value = m.input.Value()
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
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
	var statusView string
	if m.showStatus {
		switch m.status {
		case StatusSuccess:
			statusView = lipgloss.NewStyle().
				Foreground(lipgloss.Color("42")). // Green color
				Render("✓ " + m.statusMsg)
		case StatusError:
			statusView = lipgloss.NewStyle().
				Foreground(lipgloss.Color("161")). // Red color
				Render("✗ " + m.statusMsg)
		}
		statusView = "\n" + statusView
	}

	return fmt.Sprintf(
		"%s\n\n%s%s\n\n%s",
		m.prompt,
		m.input.View(),
		statusView,
		m.help.View(m.keys),
	)
}
