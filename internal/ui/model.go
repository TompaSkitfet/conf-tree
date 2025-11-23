package ui

import (
	"github.com/TompaSkitfet/conf-tree/internal/domain"
	"github.com/TompaSkitfet/conf-tree/internal/ui/components/tree"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
)

type Model struct {
	Tree   tree.Tree
	Width  int
	Height int
}

func New(root *domain.Node) Model {
	return Model{
		Tree: tree.New(root.Children),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		UpdatePanelWidths(m.Width, m.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, Keys.Up):
			m.Tree.MoveUp()
		case key.Matches(msg, Keys.Down):
			m.Tree.MoveDown()
		case key.Matches(msg, Keys.Right):
			m.Tree.MoveRight()
		case key.Matches(msg, Keys.Left):
			m.Tree.MoveLeft()
		}

	}
	return m, nil
}

func (m Model) View() string {
	selected := m.Tree.Selected()
	right := "No selection"
	if selected != nil {
		right = BuildRightTree(selected)
	}
	return TwoPanels(m.Tree.View(), right)
}
