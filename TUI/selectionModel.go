package tui

import (
	log "Attimo/logging"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	maxVisibleItems    = 10
	errorMessage       = "An error occurred: %v"
	errorReplyContinue = "Continue"
)

type selectionKeyMap struct {
	keyMap
	Enter key.Binding
	Up    key.Binding
	Down  key.Binding
}

func newSelectionKeyMap() selectionKeyMap {
	return selectionKeyMap{
		keyMap: NewKeyMap(),

		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("⏎", "confirm"),
		),

		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "move up"),
		),

		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "move down"),
		),
	}
}
func (k selectionKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.HardQuit, k.Help},
		{k.Enter, k.Up, k.Down},
	}
}

type selectionModel struct {
	tuiWindow

	keys      selectionKeyMap
	prompt    string
	userInput textinput.Model
	values    []string
	filtered  []string
	cursorPos int

	maxWidth   int // Maximum width of any string in values
	startIndex int // Start index for viewport sliding
}

func newSelectionModel(prompt string, values []string, logger *log.Logger) (*selectionModel, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	ti := textinput.New()
	ti.Placeholder = alluringString
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	// Calculate maximum width
	maxWidth := 0
	for _, value := range values {
		width := utf8.RuneCountInString(value)
		if width > maxWidth {
			maxWidth = width
		}
	}

	return &selectionModel{
		tuiWindow: tuiWindow{
			help:   help.New(),
			logger: logger,
		},

		keys:       newSelectionKeyMap(),
		prompt:     prompt,
		userInput:  ti,
		values:     values,
		filtered:   values,
		cursorPos:  0,
		startIndex: 0,
		maxWidth:   maxWidth,
	}, nil
}

func (m selectionModel) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, textinput.Blink)
}

// filterValues excludes values that do not match the input
func filterValues(values []string, input string) []string {
	if input == "" {
		return values
	}

	var filtered []string
	lowerInput := strings.ToLower(input)

	for _, value := range values {
		if strings.Contains(strings.ToLower(value), lowerInput) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func (m selectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.logger.LogInfo("Quitting TUI")
			return m, tea.Quit
		case key.Matches(msg, m.keys.Enter):
			if len(m.filtered) > 0 {
				m.logger.LogInfo("Selected item: %s", m.filtered[m.cursorPos])
				m.selected = m.cursorPos
				return m, tea.Quit
			}
			m.logger.LogInfo("No item to match for: %v", m.userInput.Value())

		case key.Matches(msg, m.keys.Up):
			m.moveUp()

		case key.Matches(msg, m.keys.Down):
			m.moveDown()

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		default:
			// Handle text input
			m.userInput, _ = m.userInput.Update(msg)
			m.filtered = filterValues(m.values, m.userInput.Value())
			m.resetCursorMaybe()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	return m, nil
}

func (m *selectionModel) moveUp() {
	if m.cursorPos > 0 {
		m.cursorPos--
		// Adjust viewport if cursor moves above visible area
		if m.cursorPos < m.startIndex {
			m.startIndex = m.cursorPos
		}
	}
}

func (m *selectionModel) moveDown() {
	if m.cursorPos < len(m.filtered)-1 {
		m.cursorPos++
		// Adjust viewport if cursor moves below visible area
		if m.cursorPos >= m.startIndex+maxVisibleItems {
			m.startIndex = m.cursorPos - maxVisibleItems + 1
		}
	}
}

func (m *selectionModel) resetCursorMaybe() {
	// Reset cursor and viewport when filter changes
	if len(m.filtered) <= m.cursorPos {
		m.cursorPos = 0
		m.startIndex = 0
	}

}

func (m selectionModel) View() string {
	style := getSingleBoxStyle(min(m.maxWidth+len(CURSOR), m.width))

	var sb strings.Builder

	// Calculate visible range
	endIndex := m.startIndex + maxVisibleItems
	if endIndex > len(m.filtered) {
		endIndex = len(m.filtered)
	}

	// Display only the visible portion of the list
	for i := m.startIndex; i < endIndex; i++ {
		// Pad the string to match maxWidth
		paddedValue := m.filtered[i] + strings.Repeat(" ", m.maxWidth-utf8.RuneCountInString(m.filtered[i]))

		if i == m.cursorPos {
			sb.WriteString(CURSOR + " " + paddedValue)
		} else {
			sb.WriteString(NOTCURSOR + " " + paddedValue)
		}

		// Add newline if not the last visible item
		if i < endIndex-1 {
			sb.WriteString("\n")
		}
	}

	// Add scrolling indicators if necessary
	view := sb.String()

	delimiter := "\n"
	if m.startIndex > 0 {
		view = UPCURSOR + delimiter + view
	} else {
		view = delimiter + view
	}
	if endIndex < len(m.filtered) {
		view = view + delimiter + DOWNCURSOR
	} else {
		view = view + delimiter
	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n%s",
		m.prompt,
		m.userInput.View(),
		style.Render(view),
		m.help.View(m.keys),
	)
}
