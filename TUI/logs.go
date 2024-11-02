package tui

import (
	log "Attimo/logging"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	maxLogs = 1000
)

type logsmodel struct {
	tuiWindow

	logs  []string
	index int
	count int

	scrollOffset int
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

func LogsModel() (*log.Logger, *logsmodel, error) {
	model := &logsmodel{}
	logger, err := log.InitLoggingWithWriter(model)
	if err != nil {
		return nil, nil, err
	}
	return logger, &logsmodel{tuiWindow: tuiWindow{logger: logger}}, nil
}

func (m logsmodel) Init() tea.Cmd {
	return nil
}

func (m logsmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			m.logger.LogInfo("Quitting TUI logs")
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.Down):
			m.logger.LogInfo("Scrolling down")
			m.scrollOffset = max(0, m.scrollOffset-1)
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Up):
			m.logger.LogInfo("Scrolling up")
			m.scrollOffset = min(m.count-m.height, m.scrollOffset+1)
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height - 1 // Reserve one line for scroll indicator
	}
	return m, nil
}

// View renders the logs
// handles wrapping, padding, and scrolling
func (m logsmodel) View() string {
	logStyle := getLogStyle()
	boxStyle := getBoxStyle(false, getFractionInt(m.width, 0.25))

	var result []string
	var allLines []string

	for _, log := range m.logs {
		wrapped := lipgloss.NewStyle().Width(m.width).Render(log)
		lines := strings.Split(wrapped, "\n")
		allLines = append(allLines, lines...)
	}

	totalLines := len(allLines)
	visibleLines := min(m.height, totalLines)

	startIndex := max(0, totalLines-visibleLines-m.scrollOffset)
	endIndex := min(totalLines, startIndex+visibleLines)

	result = allLines[startIndex:endIndex]

	scrollIndicator := fmt.Sprintf("Showing %d-%d of %d", startIndex+1, endIndex, totalLines)
	result = append(result, boxStyle.Render(scrollIndicator))

	return logStyle.Render(strings.Join(result, "\n"))
}
