package tui

import (
	"fmt"
	"strings"

	ctrl "Attimo/control"
	log "Attimo/logging"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type closedModel struct {
	tuiWindow

	keys          selectionKeyMap
	pointers      []string
	cursor        int
	startIndex    int
	selected      string
	step          closedStep
	timeHandler   *TimeHandler
	closeColumns  []string
	currentColumn int
	closeValues   map[string]string
	input         *inputModel
	control       *ctrl.Controller
}

type closedStep int

const (
	selectItem closedStep = iota
	enterTime
	completeRow
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
		closeValues: make(map[string]string),
		control:     control,
	}, nil
}

func (m *closedModel) Init() tea.Cmd {
	return nil
}

func (m *closedModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.step == completeRow && m.input != nil {
		var model tea.Model
		model, cmd = m.input.Update(msg)
		if im, ok := model.(*inputModel); ok {
			m.input = im
		}
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.step {
		case selectItem:
			return m.handleSelectItemInput(msg)
		case enterTime:
			return m.handleTimeInput(msg)
		case completeRow:
			return m.handleCompleteRowInput(msg)
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
	if key.Matches(msg, m.keys.Quit) {
		m.step = selectItem
		return m, nil
	}

	if key.Matches(msg, m.keys.Enter) {
		parts := strings.Split(m.selected, ":")
		if len(parts) != 2 {
			return m, tea.Quit
		}

		category := parts[0]

		condition := &ctrl.ColumnCondition{
			ExcludeColumn: []string{"Closed"},
			FillBehavior:  "close",
		}

		closeColumns, err := m.control.GetCategoryColumns(m.logger, category, condition)
		if err != nil {
			m.logger.LogErr("Could not get columns: %v", err)
			return m, tea.Quit
		}

		if len(closeColumns) > 0 {
			m.closeColumns = closeColumns
			m.currentColumn = 0
			// input model for columns
			input, err := newInputModel(fmt.Sprintf("Enter value for %s:", closeColumns[0]), m.logger)
			if err != nil {
				m.logger.LogErr("Could not create input model: %v", err)
				return m, tea.Quit
			}
			m.input = input
			m.input.input.Focus()
			m.step = completeRow
			cmd := m.input.Init()
			return m, tea.Batch(textinput.Blink, cmd)
		}
		return m, tea.Quit
	}
	model, cmd := m.timeHandler.Update(msg)
	if timeHandler, ok := model.(*TimeHandler); ok {
		m.timeHandler = timeHandler
	}
	return m, cmd
}

func (m *closedModel) handleCompleteRowInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Quit):
		m.step = enterTime
		return m, nil
	case key.Matches(msg, m.keys.Enter):
		// Store the current value
		m.closeValues[m.closeColumns[m.currentColumn]] = m.input.value

		// Move to next column or finish
		m.currentColumn++
		if m.currentColumn >= len(m.closeColumns) {
			return m, tea.Quit
		}

		// Create input for next column
		var err error
		m.input, err = newInputModel(fmt.Sprintf("Enter value for %s:", m.closeColumns[m.currentColumn]), m.logger)
		if err != nil {
			m.logger.LogErr("Could not create input model: %v", err)
			return m, tea.Quit
		}
		m.input.input.Focus()
		return m, tea.Batch(textinput.Blink)
	}

	// Handle regular input updates with proper type assertion
	model, cmd := m.input.Update(msg)
	if inputModel, ok := model.(*inputModel); ok {
		m.input = inputModel
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
	case completeRow:
		if len(m.closeColumns) > 0 {
			return fmt.Sprintf(
				"Completing close fields (%d/%d):\n\n%s",
				m.currentColumn+1,
				len(m.closeColumns),
				m.input.View(),
			)
		}
		return "No additional fields to complete"
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
