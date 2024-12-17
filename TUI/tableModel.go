package tui

import (
	log "Attimo/logging"
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

type tableModel struct {
	logger *log.Logger
	help   help.Model

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

func (k tableKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help}
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

func arrayToCols(cols []string) []table.Column {
	columns := make([]table.Column, len(cols))
	for i, col := range cols {
		columns[i] = table.Column{
			Title: col,
			// TODO make this dynamic
			Width: 20,
		}
	}
	return columns
}

func arrayToRows(cols []string, row []map[string]string) []table.Row {
	rows := make([]table.Row, len(row))
	for i, r := range row {
		row := make([]string, len(cols))
		for j, col := range cols {
			row[j] = fmt.Sprintf("%v", r[col])
		}
		rows[i] = row
	}
	return rows
}

func newTableModel(logger *log.Logger, cols []string, row []map[string]string) (tableModel, error) {
	if logger == nil {
		return tableModel{}, fmt.Errorf(log.LoggerNilString)
	}
	if len(cols) == 0 {
		return tableModel{}, fmt.Errorf("columns cannot be empty")
	}
	if len(row) == 0 {
		return tableModel{}, fmt.Errorf("rows cannot be empty")
	}

	columns := arrayToCols(cols)
	rows := arrayToRows(cols, row)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	m := tableModel{
		logger:   logger,
		keys:     newTableKeyMap(),
		help:     help.New(),
		cols:     columns,
		rows:     rows,
		focus:    true,
		viewport: viewport.New(0, 20),
		styles:   s,
	}

	m.viewport.Height = 20 - lipgloss.Height(m.headersView())
	m.UpdateViewport()

	return m, nil
}

func (m tableModel) Init() tea.Cmd {
	return nil
}

// Update is the Bubble Tea update loop.
func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Enter):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.MoveUp):
			m.MoveUp(1)
		case key.Matches(msg, m.keys.MoveDown):
			m.MoveDown(1)
		case key.Matches(msg, m.keys.MoveRight):
			m.MoveUp(m.viewport.Height)
		case key.Matches(msg, m.keys.MoveLeft):
			m.MoveDown(m.viewport.Height)
		case key.Matches(msg, m.keys.PageUp):
			m.MoveUp(m.viewport.Height / 2)
		case key.Matches(msg, m.keys.PageDown):
			m.MoveDown(m.viewport.Height / 2)
		case key.Matches(msg, m.keys.GotoTop):
			m.GotoTop()
		case key.Matches(msg, m.keys.GotoBottom):
			m.GotoBottom()
		}
	}

	return m, nil
}

// MoveUp moves the selection up by any number of rows.
// It can not go above the first row.
func (m *tableModel) MoveUp(n int) {
	m.cursor = clamp(m.cursor-n, 0, len(m.rows)-1)
	switch {
	case m.start == 0:
		m.viewport.SetYOffset(clamp(m.viewport.YOffset, 0, m.cursor))
	case m.start < m.viewport.Height:
		m.viewport.YOffset = (clamp(clamp(m.viewport.YOffset+n, 0, m.cursor), 0, m.viewport.Height))
	case m.viewport.YOffset >= 1:
		m.viewport.YOffset = clamp(m.viewport.YOffset+n, 1, m.viewport.Height)
	}
	m.UpdateViewport()
}

// MoveDown moves the selection down by any number of rows.
// It can not go below the last row.
func (m *tableModel) MoveDown(n int) {
	m.cursor = clamp(m.cursor+n, 0, len(m.rows)-1)
	m.UpdateViewport()

	switch {
	case m.end == len(m.rows) && m.viewport.YOffset > 0:
		m.viewport.SetYOffset(clamp(m.viewport.YOffset-n, 1, m.viewport.Height))
	case m.cursor > (m.end-m.start)/2 && m.viewport.YOffset > 0:
		m.viewport.SetYOffset(clamp(m.viewport.YOffset-n, 1, m.cursor))
	case m.viewport.YOffset > 1:
	case m.cursor > m.viewport.YOffset+m.viewport.Height-1:
		m.viewport.SetYOffset(clamp(m.viewport.YOffset+1, 0, 1))
	}
}

// GotoTop moves the selection to the first row.
func (m *tableModel) GotoTop() {
	m.MoveUp(m.cursor)
}

// GotoBottom moves the selection to the last row.
func (m *tableModel) GotoBottom() {
	m.MoveDown(len(m.rows))
}

// UpdateViewport updates the list content based on the previously defined
// columns and rows.
func (m *tableModel) UpdateViewport() {
	renderedRows := make([]string, 0, len(m.rows))

	// Render only rows from: m.cursor-m.viewport.Height to: m.cursor+m.viewport.Height
	// Constant runtime, independent of number of rows in a table.
	// Limits the number of renderedRows to a maximum of 2*m.viewport.Height
	if m.cursor >= 0 {
		m.start = clamp(m.cursor-m.viewport.Height, 0, m.cursor)
	} else {
		m.start = 0
	}
	m.end = clamp(m.cursor+m.viewport.Height, m.cursor, len(m.rows))
	for i := m.start; i < m.end; i++ {
		renderedRows = append(renderedRows, m.renderRow(i))
	}

	m.viewport.SetContent(
		lipgloss.JoinVertical(lipgloss.Left, renderedRows...),
	)
}

func (m tableModel) headersView() string {
	s := make([]string, 0, len(m.cols))
	for _, col := range m.cols {
		if col.Width <= 0 {
			continue
		}
		style := lipgloss.NewStyle().Width(col.Width).MaxWidth(col.Width).Inline(true)
		renderedCell := style.Render(runewidth.Truncate(col.Title, col.Width, "…"))
		s = append(s, m.styles.Header.Render(renderedCell))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, s...)
}

func (m *tableModel) renderRow(r int) string {
	s := make([]string, 0, len(m.cols))
	for i, value := range m.rows[r] {
		if m.cols[i].Width <= 0 {
			continue
		}
		style := lipgloss.NewStyle().Width(m.cols[i].Width).MaxWidth(m.cols[i].Width).Inline(true)
		renderedCell := m.styles.Cell.Render(style.Render(runewidth.Truncate(value, m.cols[i].Width, "…")))
		s = append(s, renderedCell)
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, s...)

	if r == m.cursor {
		return m.styles.Selected.Render(row)
	}

	return row
}

// View renders the component.
func (m tableModel) View() string {
	return m.headersView() + "\n" + m.viewport.View() + "\n" + m.help.View(m.keys)
}

func clamp(v, low, high int) int {
	return min(max(v, low), high)
}
