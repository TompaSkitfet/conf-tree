package ui

import "github.com/charmbracelet/lipgloss"

var (
	Border = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

	TopBar     = Border.Width(160).Height(3)
	LeftPanel  = Border.Width(80).Height(30)
	RightPanel = Border.Width(80).Height(30)
)

func UpdatePanelWidths(terminalWidth, terminalHeight int) {
	panelWidth := (terminalWidth / 2) - 2
	TopBar = Border.Width(terminalWidth - 3).Height(3)
	LeftPanel = Border.Width(panelWidth).Height(terminalHeight - 10)
	RightPanel = Border.Width(panelWidth).Height(terminalHeight - 10)
}
