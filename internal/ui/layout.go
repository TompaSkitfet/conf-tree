package ui

import "github.com/charmbracelet/lipgloss"

func TwoPanels(left, right string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		LeftPanel.Render(left),
		RightPanel.Render(right))
}
