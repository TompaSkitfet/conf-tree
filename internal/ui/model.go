package ui

import (
	"github.com/TompaSkitfet/conf-tree/internal/domain"
	"github.com/TompaSkitfet/conf-tree/internal/ui/components/tree"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

type Model struct {
	Tree        tree.Tree
	Width       int
	Height      int
	ShowOverlay bool
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
		case msg.String() == "o":
			m.ShowOverlay = !m.ShowOverlay
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

	base := TwoPanels(m.Tree.View(), right, BuildBreadcrumbs(selected))

	if m.ShowOverlay {
		overlayPanel := lipgloss.NewStyle().
			Width(20).Height(2).Border(lipgloss.RoundedBorder()).Align(lipgloss.Center).Render(selected.Key)
		return overlay.Composite(overlayPanel, base, overlay.Center, overlay.Center, 0, 0)

	}

	return base
}
