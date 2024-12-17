package tui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	ctrl "Attimo/control"
	log "Attimo/logging"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type closedModel struct {
	tuiWindow

	keys       selectionKeyMap
	pointers   []string
	cursor     int
	startIndex int
	selected   string
	step       closedStep
	timeInput  *inputModel
}

type closedStep int

const (
	selectItem closedStep = iota
	enterTime
)

const (
	datetimeFormat = "2006-01-02 15:04:05"
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

	timeInput, err := newInputModel("Enter close time:", logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create time input: %v", err)
	}

	timeInput.input.Focus()

	return &closedModel{
		tuiWindow: tuiWindow{
			logger: logger,
		},
		keys:      newSelectionKeyMap(),
		pointers:  pointers,
		step:      selectItem,
		timeInput: timeInput,
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

func parseTimeInput(input string) (string, error) {
	input = strings.TrimSpace(input)
	now := time.Now()

	switch {
	case input == "":
		return now.Format(datetimeFormat), nil

	case strings.HasPrefix(input, "+"):
		minutesStr := strings.TrimPrefix(input, "+")
		minutes, err := strconv.Atoi(minutesStr)
		if err != nil {
			return "", fmt.Errorf("invalid minutes format: %v", err)
		}

		return now.Add(time.Duration(minutes) * time.Minute).Format(datetimeFormat), nil

	case strings.HasPrefix(input, "-"):
		minutesStr := strings.TrimPrefix(input, "-")
		minutes, err := strconv.Atoi(minutesStr)
		if err != nil {
			return "", fmt.Errorf("invalid minutes format: %v", err)
		}

		return now.Add(-time.Duration(minutes) * time.Minute).Format(datetimeFormat), nil

	default:
		_, err := time.Parse(datetimeFormat, input)
		if err != nil {
			return "", fmt.Errorf("invalid time format: %v", err)
		}
		return input, nil
	}
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

// In handleTimeInput method
func (m *closedModel) handleTimeInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if key.Matches(msg, m.keys.Quit) {
		m.step = selectItem
		return m, nil
	}

	if key.Matches(msg, m.keys.Enter) {
		// Parse and validate the time input
		parsedTime, err := parseTimeInput(m.timeInput.input.Value())
		if err != nil {
			m.timeInput.SetStatus(StatusError, fmt.Sprintf("Invalid time format: %v", err))
			return m, nil
		}
		// Store the final value and quit
		m.timeInput.value = parsedTime
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.timeInput.input, cmd = m.timeInput.input.Update(msg) //delegate to underlying text input model
	return m, cmd
}

func (m *closedModel) View() string {
	switch m.step {
	case selectItem:
		return m.viewSelectItem()
	case enterTime:
		return fmt.Sprintf(
			"Enter close time for:\n\n%s",
			m.timeInput.View(),
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
