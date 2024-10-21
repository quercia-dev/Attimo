package tui

import (
	"strings"

	log "Attimo/logging"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	maxLogs = 100
)

type logsmodel struct {
	logs  []string
	index int
	count int

	width  int
	height int
}

func (m *logsmodel) Write(p []byte) (n int, err error) {

	logMessage := string(p)
	m.logs = append(m.logs, logMessage)
	m.index = (m.index + 1) % maxLogs
	if m.count < maxLogs {
		m.count++
	}

	return len(p), nil
}

func LogsModel() logsmodel {
	return logsmodel{}
}

func (m logsmodel) Init() tea.Cmd {
	return nil
}

func (m logsmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, DefaultKeyMap.Quit) {
			log.LogInfo("Quitting TUI logs")
			return m, tea.Quit
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}
	return m, nil
}

// View renders the logs
// handles wrapping and padding
func (m logsmodel) View() string {
	logStyle := getLogStyle()

	var result []string

	for _, log := range m.logs {
		wrapped := lipgloss.NewStyle().Width(m.width).Render(log)

		lines := strings.Split(wrapped, "\n")
		result = append(result, lines...)
	}

	totalLines := len(result)
	if totalLines < m.height {
		padding := m.height - totalLines
		for i := 0; i < padding; i++ {
			result = append(result, "") // Add empty lines for padding
		}
	}

	return logStyle.Render(strings.Join(result, "\n"))
}
