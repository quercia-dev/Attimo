package tui

import (
	"strings"

	log "Attimo/logging"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	maxLogs = 100
)

type logsmodel struct {
	logs  []string
	index int
	count int
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
		if msg.String() == hardQuitKey {
			log.LogInfo("Quitting TUI logs")
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m logsmodel) View() string {
	return strings.Join(m.logs, "\n")
}
