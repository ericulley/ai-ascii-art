package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type questionModel struct {
	asciiArt      string
	questions     []string
	questionIndex int
	choices       []string
	cursorIndex   int
	width int
	height int
}

func NewQuestionModel(art string) questionModel {
	return questionModel{
		asciiArt: art,
		questions: []string{
			"Would you like to save this art?",
			"Enter a name: ",
			"Would you like to exit or generate more art?",
		},
		questionIndex: 0,
		choices:       []string{"Yeah", "Na", "Exit", "Chat"},
		cursorIndex:   0,
	}
}

func (m questionModel) Init() tea.Cmd {
	return nil
}

func (m questionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {
		// These keys should exit the program.
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursorIndex > 0 {
				m.cursorIndex--
			}
		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursorIndex < 1 {
				m.cursorIndex++
			}
		// The "enter" key selects the state for the item that the cursor is pointing at.
		case "enter":
			switch m.questionIndex {
			case 0: // "Would you like to save this art?"
				if m.cursorIndex == 0 {
					return NewPromptModel(m.asciiArt).Update(msg)
				} else if m.cursorIndex == 1 {
					m.questionIndex = 2
					return m, nil
				}
			case 1: // "Enter a name: "
				// save in the db
	
			case 2: // "Would you like to exit or generate more art?"
				if m.cursorIndex == 0 {
					return m, tea.Quit
				} else if m.cursorIndex == 1 {
					return NewChatModel().Update(msg);
				}
			}

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m questionModel) View() string {
	var s string
	// Display ascii art
	if m.asciiArt != "" {
		s = fmt.Sprintf(m.asciiArt + "\n\n")
	}
	// Display the prompt
	s = s + m.questions[m.questionIndex] + "\n"

	// Iterate over our choices for the first two questions
	if m.questionIndex < 2 {
		for i, choice := range m.choices[:2] {
			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursorIndex == i {
				cursor = ">" // cursor!
			}
			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	} else if m.questionIndex == 2 {
		for i, choice := range m.choices[2:] {
			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursorIndex == i {
				cursor = ">" // cursor!
			}
			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	}

	// Send the UI for rendering
	return s
}
