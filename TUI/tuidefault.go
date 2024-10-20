package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

const (
	hardQuitKey = "ctrl+c"
	CURSOR      = "❯❯❯"
	NOTCURSOR   = "   "
)

type KeyMap struct {
	Quit key.Binding

	Enter key.Binding

	Up    key.Binding
	Down  key.Binding
	Right key.Binding
	Left  key.Binding
}

type CustomKeyMap struct {
	Quit        key.Binding
	Enter       key.Binding
	GreedyEnter key.Binding
	Up          key.Binding
	Down        key.Binding
	Right       key.Binding
	Left        key.Binding
}

var DefaultKeyMap = CustomKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q/esc", "quit"),
	),

	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm"),
	),

	GreedyEnter: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter/space", "confirm"),
	),

	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("→/l", "move right"),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("←/h", "move left"),
	),
}

func getBoxStyle(selected bool, width int) lipgloss.Style {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 1).
		AlignHorizontal(lipgloss.Center).
		Width(width)

	if selected {
		return style.Background(lipgloss.Color("#205c63"))
	}
	return style
}

func getLogStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cd00cd"))
}
