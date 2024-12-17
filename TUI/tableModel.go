package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
)

type tableModel struct {
	tuiWindow

	keys     tableKeyMap
	viewport viewport.Model

	cols []table.Column
	rows []table.Row

	cursor int
	focus  bool
	styles table.Styles

	start int
	end   int
}

type tableKeyMap struct {
	keyMap
	Enter      key.Binding
	MoveRight  key.Binding
	MoveLeft   key.Binding
	MoveUp     key.Binding
	MoveDown   key.Binding
	PageUp     key.Binding
	PageDown   key.Binding
	GotoTop    key.Binding
	GotoBottom key.Binding
}

func (k tableKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.HardQuit, k.Help},
		{k.Enter, k.MoveRight, k.MoveLeft, k.MoveUp, k.MoveDown},
	}
}

func newTableKeyMap() tableKeyMap {
	keys := NewKeyMap()
	keys.Help = key.NewBinding(
		key.WithKeys("ctrl+h", "?"),
		key.WithHelp("ctrl+h/?", "help"),
	)

	return tableKeyMap{
		keyMap: keys,

		Enter: key.NewBinding(
			key.WithKeys("enter", " "),
			key.WithHelp("⏎/' '", "confirm"),
		),

		MoveRight: key.NewBinding(
			key.WithKeys("l", "right"),
			key.WithHelp("→/l", "move right"),
		),
		MoveLeft: key.NewBinding(
			key.WithKeys("h", "left"),
			key.WithHelp("←/h", "move left"),
		),
		MoveUp: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("↑/k", "move up"),
		),
		MoveDown: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("↓/j", "move down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("ctrl+b", "pageup"),
			key.WithHelp("ctrl+b/pageup", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("ctrl+f", "pagedown"),
			key.WithHelp("ctrl+f/pagedown", "page down"),
		),
		GotoTop: key.NewBinding(
			key.WithKeys("g", "home"),
			key.WithHelp("g/home", "go to top"),
		),
		GotoBottom: key.NewBinding(
			key.WithKeys("G", "end"),
			key.WithHelp("G/end", "go to bottom"),
		),
	}

}
