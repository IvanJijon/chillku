package ui

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

type Countdown struct {
	remainingTime time.Time
	style         lipgloss.Style
}

func InitCountdown() *Countdown {
	return &Countdown{
		remainingTime: time.Time{},
		style:         lipgloss.Style{},
	}
}

func (cd *Countdown) View() string {
	return cd.style.Render(cd.remainingTime.String())
}
