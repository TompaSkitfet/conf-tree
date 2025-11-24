package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit  key.Binding
	Save  key.Binding
	Up    key.Binding
	Down  key.Binding
	Right key.Binding
	Left  key.Binding
}

var Keys = KeyMap{
	Quit:  key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	Save:  key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "save")),
	Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
	Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
	Right: key.NewBinding(key.WithKeys("right", "l", "enter"), key.WithHelp("→/l/enter", "open/edit")),
	Left:  key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "back")),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Save, k.Up, k.Down, k.Right, k.Left}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{k.Save, k.Quit},
	}
}
