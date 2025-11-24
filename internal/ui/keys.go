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
	Save:  key.NewBinding(key.WithKeys("s")),
	Up:    key.NewBinding(key.WithKeys("up", "k")),
	Down:  key.NewBinding(key.WithKeys("down", "j")),
	Right: key.NewBinding(key.WithKeys("right", "l", "enter")),
	Left:  key.NewBinding(key.WithKeys("left", "h")),
}
