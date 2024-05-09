package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/ellielle/bootdev-stats/internal/cache"
)

// FIXME: revisit this and determine if it's better
// to send it on its own

// NOTE: client may be rewritten as TUI is implemented
// TODO: GetArchmages

// Send the cache along for the ride
type Client struct {
	cache cache.Cache
}

// Bubble Tea

// model holds the application's state
type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	cache    cache.Cache
}

// initialModel initializes the state for Bubble Tea
func initialModel() model {
	return model{
		choices:  []string{"archmage", "stats", "daily", "karma"},
		selected: make(map[int]struct{}),
		cache:    cache.NewCache(300 * time.Second),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Handle key press events
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""
	for i, choice := range m.choices {

		// if the cursor hasn't moved yet, it will be invisible
		// otherwise, change the cursor to `>` and move it to that choice
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">"
		}

		// as with above, a selection is initially unchecked
		// change the mark by the selection to `x` if it's checked
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		// render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)

	}

	return s
}
