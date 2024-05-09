package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// create a new Bubble Tea program
	// exit the program with a non-zero exit code
	// if something is amiss
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Printf("uh oh, it broke! %v", err.Error())
		os.Exit(1)
	}
}
