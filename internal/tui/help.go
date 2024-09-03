package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Delete key.Binding
	NextPage key.Binding
	PrevPage key.Binding
	Save key.Binding
	New key.Binding
	Enter key.Binding
	Up    key.Binding
	Down  key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Enter, k.Up, k.Down, k.Save, k.New, k.NextPage, k.PrevPage, k.Delete},
		{k.Help, k.Quit},
	}
}

var keys = keyMap{
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete saved word"),
	),
	NextPage: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next page"),
	),
	PrevPage: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "prev page"),
	),
	New: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "Refresh random words"),
	),
	Save: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "Save a word"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Proceed"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}


