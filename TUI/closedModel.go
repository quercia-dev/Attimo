package tui

import (
	"fmt"
	"strings"

	ctrl "Attimo/control"
	log "Attimo/logging"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type closedModel struct {
	tuiWindow

	keys        selectionKeyMap
	pointers    []string
	cursor      int
	startIndex  int
	selected    string
	step        closedStep
	timeHandler *TimeHandler
}

type closedStep int

const (
	selectItem closedStep = iota
	enterTime
)

func newClosedModel(logger *log.Logger, control *ctrl.Controller) (*closedModel, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	// Get pending items
	pointers, err := control.GetPendingPointers(logger)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending pointers: %v", err)
	}

	if len(pointers) == 0 {
		return nil, fmt.Errorf("no pending items to close")
	}

	timeHandler, err := NewTimeHandler("Enter close time:", logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create time handler: %v", err)
	}

	return &closedModel{
		tuiWindow: tuiWindow{
			logger: logger,
		},
		keys:        newSelectionKeyMap(),
		pointers:    pointers,
		step:        selectItem,
		timeHandler: timeHandler,
	}, nil
}

func (m *closedModel) Init() tea.Cmd {
	return nil
}

func (m *closedModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.step {
		case selectItem:
			return m.handleSelectItemInput(msg)
		case enterTime:
			return m.handleTimeInput(msg)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	return m, nil
}

func (m *closedModel) handleSelectItemInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Quit):
		m.logger.LogInfo("Quitting close selection")
		return m, tea.Quit

	case key.Matches(msg, m.keys.Up):
		if m.cursor > 0 {
			m.cursor--
			if m.cursor < m.startIndex {
				m.startIndex = m.cursor
			}
		}
		return m, nil

	case key.Matches(msg, m.keys.Down):
		if m.cursor < len(m.pointers)-1 {
			m.cursor++
			if m.cursor >= m.startIndex+maxVisibleItems {
				m.startIndex = m.cursor - maxVisibleItems + 1
			}
		}
		return m, nil

	case key.Matches(msg, m.keys.Enter):
		m.selected = m.pointers[m.cursor]
		m.step = enterTime
		return m, nil
	}

	return m, nil
}

func (m *closedModel) handleTimeInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	model, cmd := m.timeHandler.Update(msg)
	if timeHandler, ok := model.(*TimeHandler); ok {
		m.timeHandler = timeHandler
	}
	return m, cmd
}

func (m *closedModel) View() string {
	switch m.step {
	case selectItem:
		return m.viewSelectItem()
	case enterTime:
		return fmt.Sprintf(
			"Enter close time for:\n\n%s",
			m.timeHandler.View(),
		)
	default:
		return "Unknown step"
	}
}

func (m *closedModel) viewSelectItem() string {
	var sb strings.Builder
	sb.WriteString("Select item to close:\n\n")

	// Calculate visible range
	endIndex := m.startIndex + maxVisibleItems
	if endIndex > len(m.pointers) {
		endIndex = len(m.pointers)
	}

	// Display visible items
	for i := m.startIndex; i < endIndex; i++ {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		sb.WriteString(fmt.Sprintf("%s %s\n", cursor, m.pointers[i]))
	}

	// Add scroll indicators
	if m.startIndex > 0 {
		sb.WriteString("\n" + UPCURSOR)
	}
	if endIndex < len(m.pointers) {
		sb.WriteString("\n" + DOWNCURSOR)
	}

	sb.WriteString("\n\n(↑/↓) navigate • (enter) select • (esc) quit")
	return sb.String()
}
