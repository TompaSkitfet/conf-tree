package main

import (
	"fmt"
	"os"

	"github.com/TompaSkitfet/conf-tree/internal/config"
	"github.com/TompaSkitfet/conf-tree/internal/domain"
	"github.com/TompaSkitfet/conf-tree/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Specify a json file 'ctree <file.json>'")
		os.Exit(1)
	}

	filepath := os.Args[1]
	fileData := domain.FileData{Name: filepath, FileType: "json"}

	root, err := config.LoadJSON(filepath)
	if err != nil {
		fmt.Println("Specify a json file 'ctree <file.json>'")
		os.Exit(1)
	}

	p := tea.NewProgram(ui.New(root, fileData), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}

}
