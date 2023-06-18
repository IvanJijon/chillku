package ui

import "github.com/charmbracelet/lipgloss"

// Header is the app's ui header
type Header struct {
	content string
	style   lipgloss.Style
}

// InitHeader creates the header of the app
func InitHeader() *Header {
	s := lipgloss.NewStyle().
		Bold(true).
		Padding(1).
		AlignHorizontal(lipgloss.Center).
		Width(applicationWidth)

	return &Header{
		content: "\U0001f345 Chillku Timer\n",
		style:   s,
	}
}

// View renders the header
func (h *Header) View() string {
	return h.style.Render(h.content)
}
