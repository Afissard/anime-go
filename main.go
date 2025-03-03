package main

import (
	"anime-go/app"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	p := tea.NewProgram(app.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
