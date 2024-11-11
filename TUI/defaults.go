package tui

import (
	log "Attimo/logging"
	"math"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

const (
	hardQuitKey    = "ctrl+c"
	CURSOR         = "❯❯❯"
	NOTCURSOR      = "   "
	alluringString = "...? 🤔"
	DOWNCURSOR     = "▼ ▼ ▼"
	UPCURSOR       = "▲ ▲ ▲"

	quitMessage = "Quitting TUI"
	TUIerror    = "Error running TUI: %v"
)

type tuiWindow struct {
	logger *log.Logger
	keys   keyMap
	help   help.Model

	width  int
	height int

	selected interface{}
}

type keyMap struct {
	// Quit is the key binding for quitting the program.
	// Bound to a greater number of keys to make it easier to quit.
	Quit key.Binding
	// HardQuit is the key binding for quitting the program.
	// Is is a subset of the Quit key binding.
	// It is used for when the user is in a state where they need
	// to use the Quit key binding.
	HardQuit    key.Binding
	Enter       key.Binding
	GreedyEnter key.Binding
	Up          key.Binding
	Down        key.Binding
	Right       key.Binding
	Left        key.Binding
	Help        key.Binding
}

var DefaultKeyMap = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", hardQuitKey),
		key.WithHelp("q/esc", "quit"),
	),

	HardQuit: key.NewBinding(
		key.WithKeys(hardQuitKey),
		key.WithHelp(hardQuitKey, "quit"),
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
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},          // first column
		{k.Enter, k.GreedyEnter, k.Quit, k.Help}, // second column
	}
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

func getSingleBoxStyle(width int) lipgloss.Style {
	style := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("63")).
		AlignHorizontal(lipgloss.Center).
		Width(width)

	return style
}

func getLogStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cd00cd"))
}

func getFractionInt(width int, fraction float32) int {
	if width < 0 {
		width = 0
	}

	if fraction < 0 {
		fraction = 0
	}
	if fraction > 1 {
		fraction = 1
	}

	return int(math.Round(float64(width) * float64(fraction)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
