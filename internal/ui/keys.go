package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit  key.Binding
	Up    key.Binding
	Down  key.Binding
	Right key.Binding
	Left  key.Binding
}

var Keys = KeyMap{
	Quit:  key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	Up:    key.NewBinding(key.WithKeys("up", "k")),
	Down:  key.NewBinding(key.WithKeys("down", "j")),
	Right: key.NewBinding(key.WithKeys("right", "l")),
	Left:  key.NewBinding(key.WithKeys("left", "h")),
}
