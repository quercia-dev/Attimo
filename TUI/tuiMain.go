package tui

import (
	log "Attimo/logging"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	openItem   = "OPEN"
	closeItem  = "CLOSE"
	agendaItem = "AGENDA"
	editItem   = "EDIT"

	openShortcut   = "o"
	closeShortcut  = "c"
	agendaShortcut = "a"
	editShortcut   = "e"
)

type mainMenu struct {
	menuItems []string
	cursor    int
	shortcuts map[string]string
}

func MainModel() mainMenu {
	return mainMenu{
		menuItems: []string{openItem, closeItem, agendaItem, editItem},
		shortcuts: map[string]string{
			openShortcut:   openItem,
			closeShortcut:  closeItem,
			agendaShortcut: agendaItem,
			editShortcut:   editItem,
		},
	}
}

func (m mainMenu) Init() tea.Cmd {
	return nil
}

func (m mainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			// TODO: Handle quitting through control
			log.LogInfo("Quitting TUI")
			return m, tea.Quit

		case key.Matches(msg, DefaultKeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			}
			log.LogInfo("Cursor up")
		case key.Matches(msg, DefaultKeyMap.Down):
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}
			log.LogInfo("Cursor down")
		case key.Matches(msg, DefaultKeyMap.Enter):
			// Here you would handle the selection
			log.LogInfo("Selected item: %s", m.menuItems[m.cursor])
			return m, tea.Quit
		default:
			log.LogInfo("Key press: %s", msg.String())
			if _, exists := m.shortcuts[msg.String()]; exists {
				// Here you would handle the selection
				log.LogInfo("Selected shortcut item: %s", m.shortcuts[msg.String()])
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m mainMenu) View() string {
	s := "\n\n\n"

	style := getBoxStyle()

	// Menu items
	for i, item := range m.menuItems {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		n := 20
		// Add number prefix and format menu item
		menuText := fmt.Sprintf("%-*s", n, fmt.Sprintf("%s %d. %s", cursor, i+1, item))

		// Add shortcut if exists
		if shortcut, exists := m.shortcuts[item]; exists {
			menuText = fmt.Sprintf("%-40s %s", menuText, shortcut)
		}

		s += fmt.Sprintf("%s\n", style.Render(menuText))
	}

	// Quick open note
	s += "\nQuick open with a system shortcut\n"
	s += "\nPress q to quit\n"

	return s
}
