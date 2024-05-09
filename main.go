package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// create a new Bubble Tea program
	// exit the program with a non-zero exit code
	// if something is amiss
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf("uh oh, it broke! %v", err.Error())
	}
}
