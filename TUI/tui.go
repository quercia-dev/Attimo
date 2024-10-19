package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	menuItems []string
	cursor    int
	shortcuts map[string]string
}

func InitialModel() model {
	return model{
		menuItems: []string{"OPEN", "CLOSE", "AGENDA", "EDIT"},
		shortcuts: map[string]string{
			"CLOSE":  "#",
			"AGENDA": "#",
			"EDIT":   "#",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}
		case "enter", " ":
			// Here you would handle the selection
			return m, tea.Quit
		case "s":
			// Handle settings
			return m, nil
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "\n  INITIAL SCREEN / INITIAL BUTTONS\n\n"

	// Settings button in top right
	s += strings.Repeat(" ", 60) + "[S] SETTINGS\n\n"

	// Menu items
	for i, item := range m.menuItems {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// Add number prefix and format menu item
		menuText := fmt.Sprintf("%s %d. %s", cursor, i+1, item)

		// Add shortcut if exists
		if shortcut, exists := m.shortcuts[item]; exists {
			menuText = fmt.Sprintf("%-40s %s", menuText, shortcut)
		}

		s += fmt.Sprintf("  %s\n", menuText)
	}

	// Quick open note
	s += "\nQuick open with a system shortcut\n"
	s += "\nPress q to quit\n"

	return s
}
