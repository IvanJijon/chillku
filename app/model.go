package app

import (
	"fmt"
	"time"

	"github.com/IvanJijon/chillku/hourglass"
	"github.com/IvanJijon/chillku/ui"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model is the main application model
type Model struct {
	spinner spinner.Model
	hg      *hourglass.Hourglass
	ui      ui.UI
}

// InitialModel instantiates model at its starting state when launching the app
func InitialModel() Model {
	ui := ui.NewUI()
	return Model{
		spinner: initSpinner(),
		hg:      hourglass.NewHourglass(5 * time.Second),
		ui:      ui,
	}
}

func initSpinner() spinner.Model {

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

// Init is the app's Init method
func (m Model) Init() tea.Cmd {
	return nil
}

// Update is the app's Update method
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "q":
			return m, tea.Quit
		case "s":
			m.hg.Start()
			return m, m.spinner.Tick
		default:
			return m, nil
		}
	default:
		//TODO: handle hourglass expiration
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

// View renders the app's UI
func (m Model) View() string {
	s := m.ui.Header.View()
	s += fmt.Sprintf("\n%s\n", m.spinner.View())
	s += fmt.Sprint(m.hg.Countdown().Truncate(time.Second).String())
	s += "\n\n  press q to Quit\n"
	return s
}
