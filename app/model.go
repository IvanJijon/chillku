package app

import (
	"fmt"
	"time"

	"github.com/IvanJijon/chillku/hourglass"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	spinner spinner.Model
	hg      *hourglass.Hourglass
}

func InitialModel() model {
	// fmt.Println("model created, hourglass created")
	return model{spinner: initSpinner(), hg: hourglass.NewHourglass(5 * time.Second)}
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

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {

	s := "Header\n"
	s += fmt.Sprintf("%s \n\n", m.spinner.View())
	s += fmt.Sprint((time.Second + m.hg.CountDown()).Truncate(time.Second).String())
	s += "\n\n  press q to Quit\n"
	return s
}
