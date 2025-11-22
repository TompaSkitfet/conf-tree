package main

import (
	"fmt"

	"github.com/TompaSkitfet/conf-tree/internal/config"
	"github.com/TompaSkitfet/conf-tree/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	root, err := config.LoadJSON()
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(ui.New(root))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}

}
