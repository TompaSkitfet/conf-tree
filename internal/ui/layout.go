package ui

import "github.com/charmbracelet/lipgloss"

func TwoPanels(left, right string) string {
	panels := lipgloss.JoinHorizontal(lipgloss.Top,
		LeftPanel.Render(left),
		RightPanel.Render(right))

	layout := lipgloss.JoinVertical(lipgloss.Left, TopBar.Render("Json explorer"), panels)
	return layout

}
