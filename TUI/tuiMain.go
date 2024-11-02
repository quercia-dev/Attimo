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

type boxMenu struct {
	tuiWindow
	menuItems []string
	cursor    int
	shortcuts map[string]int
}

func boxModel(logger *log.Logger, menuItems []string, shortcuts map[string]int) (boxMenu, error) {
	if logger == nil {
		return boxMenu{}, fmt.Errorf(log.LoggerNilString)
	}
	return boxMenu{
		menuItems: menuItems,
		shortcuts: shortcuts,
	}, nil
}

func (m boxMenu) Init() tea.Cmd {
	return nil
}

func (m boxMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			m.logger.LogInfo(quitMessage + " from main menu")
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Down):
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap.GreedyEnter):
			m.selected = m.menuItems[m.cursor]
			return m, tea.Quit
		default:
			index := m.shortcuts[msg.String()]
			if index < 0 || index >= len(m.menuItems) {
				m.logger.LogWarn("Unidentified key pressed: %s", msg.String())
				return m, nil
			}
			m.selected = m.menuItems[index]
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}
	return m, nil
}

func (m boxMenu) View() string {
	s := ""

	boxWidth := getFractionInt(m.width, 0.25)
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
