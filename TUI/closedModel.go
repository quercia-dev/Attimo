package tui

import (
	"fmt"
	"strings"

	ctrl "Attimo/control"
	"Attimo/database"
	log "Attimo/logging"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// closedModel:
type closedModel struct {
	tuiWindow
	pendingItems []database.RowData
	cursor       int
	startIndex   int
	selected     database.RowData
	step         closedStep
	dateInput    *inputModel
}

type closedStep int

const (
	selectItem closedStep = iota
	enterDate
	confirmClose
)

func newClosedModel(logger *log.Logger, control *ctrl.Controller) (*closedModel, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	// Get pending items
	pendingItems, err := control.GetPendingItems(logger)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending items: %v", err)
	}

	if len(pendingItems) == 0 {
		return nil, fmt.Errorf("no pending items to close")
	}

	dateInput, err := newInputModel("Enter close date (DD-MM-YYYY):", logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create date input: %v", err)
	}

	return &closedModel{
		tuiWindow: tuiWindow{
			logger: logger,
		},
		pendingItems: pendingItems,
		step:         selectItem,
		dateInput:    dateInput,
	}, nil
}

func (m closedModel) Init() tea.Cmd {
	return nil
}

func (m closedModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.step {
		case selectItem:
			return m.handleSelectItemInput(msg)
		case enterDate:
			return m.handleDateInput(msg)
		case confirmClose:
			return m.handleConfirmInput(msg)
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	return m, nil
}

func (m closedModel) handleSelectItemInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, DefaultKeyMap.Quit):
		m.logger.LogInfo("Quitting close selection")
		return m, tea.Quit

	case key.Matches(msg, DefaultKeyMap.Up):
		if m.cursor > 0 {
			m.cursor--
			if m.cursor < m.startIndex {
				m.startIndex = m.cursor
			}
		}
		return m, nil

	case key.Matches(msg, DefaultKeyMap.Down):
		if m.cursor < len(m.pendingItems)-1 {
			m.cursor++
			if m.cursor >= m.startIndex+maxVisibleItems {
				m.startIndex = m.cursor - maxVisibleItems + 1
			}
		}
		return m, nil

	case key.Matches(msg, DefaultKeyMap.Enter):
		m.selected = m.pendingItems[m.cursor]
		m.step = enterDate
		return m, nil
	}

	return m, nil
}

func (m closedModel) handleDateInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if key.Matches(msg, DefaultKeyMap.Quit) {
		m.step = selectItem
		return m, nil
	}

	var cmd tea.Cmd
	dateInputModel, cmd := m.dateInput.Update(msg)
	m.dateInput = dateInputModel.(*inputModel)

	if m.dateInput.value != "" {
		// Date was entered, move to confirmation
		m.step = confirmClose
	}

	return m, cmd
}

func (m closedModel) handleConfirmInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, DefaultKeyMap.Enter):
		m.logger.LogInfo("Confirming close of item")
		return m, tea.Quit

	case key.Matches(msg, DefaultKeyMap.Quit):
		m.step = enterDate
		m.dateInput.value = ""
		return m, nil
	}

	return m, nil
}

func (m closedModel) View() string {
	switch m.step {
	case selectItem:
		return m.viewSelectItem()
	case enterDate:
		return m.dateInput.View()
	case confirmClose:
		return m.viewConfirm()
	default:
		return "Unknown step"
	}
}

func (m closedModel) viewSelectItem() string {
	var sb strings.Builder
	sb.WriteString("Select item to close:\n\n")

	// Calculate visible range
	endIndex := m.startIndex + maxVisibleItems
	if endIndex > len(m.pendingItems) {
		endIndex = len(m.pendingItems)
	}

	// Display visible items
	for i := m.startIndex; i < endIndex; i++ {
		item := m.pendingItems[i]
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		// Format item display
		displayText := fmt.Sprintf("%s %s - %s",
			cursor,
			item["Project"],
			item["Note"],
		)

		if createdAt, ok := item["created_at"].(string); ok {
			displayText += fmt.Sprintf(" (Opened: %s)", createdAt)
		}

		sb.WriteString(displayText + "\n")
	}

	// Add scroll indicators
	if m.startIndex > 0 {
		sb.WriteString("\n" + UPCURSOR)
	}
	if endIndex < len(m.pendingItems) {
		sb.WriteString("\n" + DOWNCURSOR)
	}

	sb.WriteString("\n\n(↑/↓) navigate • (enter) select • (esc) quit")
	return sb.String()
}

func (m closedModel) viewConfirm() string {
	var sb strings.Builder
	sb.WriteString("Confirm closing the following item:\n\n")

	sb.WriteString(fmt.Sprintf("Project: %v\n", m.selected["Project"]))
	sb.WriteString(fmt.Sprintf("Note: %v\n", m.selected["Note"]))
	sb.WriteString(fmt.Sprintf("Close Date: %s\n\n", m.dateInput.value))

	sb.WriteString("(enter) confirm • (esc) back")
	return sb.String()
}
