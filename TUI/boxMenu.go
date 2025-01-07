package tui

import (
	log "Attimo/logging"
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	openItem   = "OPEN"
	closeItem  = "CLOSE"
	agendaItem = "AGENDA"
	editItem   = "EDIT"
	logItem    = "LOGS"
)

type boxMenu struct {
	tuiWindow
	keys      menuKeyMap
	menuItems []string
	cursor    int
}

func newBoxModel(logger *log.Logger, menuItems []string) (boxMenu, error) {
	if logger == nil {
		return boxMenu{}, fmt.Errorf(log.LoggerNilString)
	}
	return boxMenu{
		tuiWindow: tuiWindow{
			help:   help.New(),
			logger: logger,
		},
		keys:      newMenuKeyMap(),
		menuItems: menuItems,
	}, nil
}

func (m boxMenu) Init() tea.Cmd {
	return nil
}

func (m boxMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.logger.LogInfo(quitMessage + " from main menu")
			return m, tea.Quit
		case key.Matches(msg, m.keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}
			return m, nil
		case key.Matches(msg, m.keys.Down):
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}
			return m, nil
		case key.Matches(msg, m.keys.Enter):
			m.selected = m.cursor
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
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
	helpView := m.help.View(m.keys)
	return s + "\n" + helpView
}
