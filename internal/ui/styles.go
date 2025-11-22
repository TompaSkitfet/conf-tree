package ui

import "github.com/charmbracelet/lipgloss"

var (
	Border = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1)

	LeftPanel  = Border.Width(40)
	RightPanel = Border
)
