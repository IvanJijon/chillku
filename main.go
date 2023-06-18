package main

import (
	"fmt"
	"os"

	"github.com/IvanJijon/chillku/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(app.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops! There's been an error: %v", err)
		os.Exit(1)
	}
}
