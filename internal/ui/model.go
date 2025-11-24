package ui

import (
	"fmt"

	"github.com/TompaSkitfet/conf-tree/internal/domain"
	"github.com/TompaSkitfet/conf-tree/internal/ui/components/tree"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

type Model struct {
	Tree        tree.Tree
	Width       int
	Height      int
	ShowOverlay bool
	Input       textinput.Model
}

func New(root *domain.Node) Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 30
	return Model{
		Tree:  tree.New(root.Children),
		Input: ti,
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	current := m.Tree.Current

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		UpdatePanelWidths(m.Width, m.Height)
	case tea.KeyMsg:
		switch {
		case m.ShowOverlay:
			var cmd tea.Cmd
			m.Input, cmd = m.Input.Update(msg)
			if msg.String() == "esc" {
				m.ShowOverlay = false
			}
			if msg.String() == "enter" {
				m.ShowOverlay = false
			}
			return m, cmd
		case key.Matches(msg, Keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, Keys.Up):
			m.Tree.MoveUp()
		case key.Matches(msg, Keys.Down):
			m.Tree.MoveDown()
		case key.Matches(msg, Keys.Right):
			if current.Type != domain.ValueNode {
				m.Tree.MoveRight()
			} else if m.Tree.Current.Type == domain.ValueNode {
				m.Input.SetValue(fmt.Sprintf("%v", m.Tree.Current.Value))
				m.ShowOverlay = !m.ShowOverlay
			}
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

	base := TwoPanels(m.Tree.View(), right, BuildBreadcrumbs(selected))

	if m.ShowOverlay {
		overlayPanel := lipgloss.NewStyle().
			Width(20).Height(2).Border(lipgloss.RoundedBorder()).Align(lipgloss.Center).Render(m.Input.View())
		return overlay.Composite(overlayPanel, base, overlay.Center, overlay.Center, 0, 0)

	}

	return base
}
