package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	db "github.com/ericulley/ascii/data"
)

type promptModel struct {
	asciiArt   string
	prompts []string
	promptIndex int
	answerField textinput.Model
	width int
}

func (m promptModel) Init() tea.Cmd {
	return textinput.Blink
}

func NewPromptModel(art string) *promptModel {
	answerField := textinput.New()
	answerField.Placeholder = "Your answer here"
	answerField.Focus()
	answerField.Width = 128
	return &promptModel{
		asciiArt: art,
		prompts: []string{"Enter a name to store this art: ", "Success! Your art was stored under "}, 
		promptIndex: 0,
		answerField: answerField, 
		width: 80,
	}
}

func (m promptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.promptIndex == 1 {
		return m, tea.Quit
	}
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.answerField.Value() != "" {
				// update database & return success message
				m.promptIndex = 1
				db.SaveArtToDB(db.AsciiRecord{Name: m.answerField.Value(), Art: m.asciiArt})
			}
			return m, nil
		}
	}
	m.answerField, cmd = m.answerField.Update(msg)
	return m, cmd
}

func (m promptModel) View() string {
	if m.width == 0 {
		return "loading..."
	}
	var prompt string
	if m.promptIndex == 1 {
		prompt = m.prompts[m.promptIndex] + m.answerField.Value()
	} else {
		prompt = m.prompts[m.promptIndex]
	}
	return lipgloss.JoinVertical(lipgloss.Left, prompt, m.answerField.View())
}