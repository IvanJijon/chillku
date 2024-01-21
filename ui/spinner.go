package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

func InitSpinner() spinner.Model {
	s := spinner.New()

	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#C15C5C")).
		PaddingTop(0).
		Align(lipgloss.Center).
		Width(30)

	return s
}
