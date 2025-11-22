package ui

import "github.com/charmbracelet/lipgloss"

var (
	Border = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1)

	TopBar     = Border.Width(160).Height(3)
	LeftPanel  = Border.Width(80).Height(30)
	RightPanel = Border.Width(80).Height(30)
)
