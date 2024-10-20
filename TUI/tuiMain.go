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
	logItem    = "LOGS"

	openShortcut   = "o"
	closeShortcut  = "c"
	agendaShortcut = "a"
	editShortcut   = "e"
	logShortcut    = "l"
)

type mainMenu struct {
	menuItems []string
	cursor    int
	shortcuts map[string]string

	width  int
	height int
}

func MainModel() mainMenu {
	return mainMenu{
		menuItems: []string{openItem, closeItem, agendaItem, editItem, logItem},
		shortcuts: map[string]string{
			openShortcut:   openItem,
			closeShortcut:  closeItem,
			agendaShortcut: agendaItem,
			editShortcut:   editItem,
			logShortcut:    logItem,
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
			log.LogInfo("Quitting TUI")
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			}
			log.LogInfo("Cursor up")
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Down):
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}
			log.LogInfo("Cursor down")
			return m, nil
		case key.Matches(msg, DefaultKeyMap.GreedyEnter):
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
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}
	return m, nil
}

func (m mainMenu) View() string {
	s := ""

	boxWidth := m.getBoxWidth()
	style := getBoxStyle(false, boxWidth)
	styleSel := getBoxStyle(true, boxWidth)

	// Menu items
	for i, item := range m.menuItems {
		if m.cursor == i {
			s += fmt.Sprintf("%s\n", styleSel.Render(item))
			continue
		}
		s += fmt.Sprintf("%s\n", style.Render(item))
	}
	return s
}

func (m mainMenu) getBoxWidth() int {
	return m.width / 4
}
