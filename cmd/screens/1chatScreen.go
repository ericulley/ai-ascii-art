package screens

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sashabaranov/go-openai"
)

type chatModel struct {
	textarea    textarea.Model
	viewport    viewport.Model
	messages    []string
	senderStyle lipgloss.Style
	err         error
	aiClient    *openai.Client
	ascii       *ascii
}

type ascii struct {
	art string
}

type asciiMsg bool

func NewChatModel() chatModel {
	ta := textarea.New()
	ta.Placeholder = "Send a message...(esc to exit)"
	ta.Focus()

	ta.Prompt = "> "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(1)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 10)
	vp.SetContent(`Ask ChatGPT to create some ascii art!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return chatModel{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		aiClient:    openai.NewClient(os.Getenv("OPENAI_API_KEY")),
		ascii:       nil,
	}
}

func (m chatModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m chatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case asciiMsg:
		return NewQuestionModel(m.ascii.art).Update(msg)
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.textarea.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			// Quit.
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case "enter":
			v := m.textarea.Value()

			if v == "" {
				// Don't send empty messages.
				return m, nil
			}

			// Simulate sending a message. In your application you'll want to
			// also return a custom command to send the message off to
			// a server.
			// Send message to openai
			resp, err := m.SendMessage(v)
			if err != nil {
				fmt.Printf("Completion error: %v\n", err)
			}
			respContent := resp.Message.Content

			m.messages = append(m.messages, m.senderStyle.Render("You: ")+v)
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.messages = append(m.messages, m.senderStyle.Render("ChatGPT: "+respContent))
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()

			// Check for ascii art code snippet and prompt to save it
			hasCodeSnippet := strings.Contains(respContent, "```")
			if hasCodeSnippet {
				start := strings.Index(respContent, "```")
				end := strings.LastIndex(respContent, "```") + 3
				m.ascii = &ascii{art: respContent[start:end]}
				return m, storedAsciiArt
			}

			return m, nil
		case tea.KeyUp.String():
			m.viewport.LineUp(1)
			return m, nil
		case tea.KeyDown.String():
			m.viewport.LineDown(1)
			return m, nil
		default:
			// Send all other keypresses to the textarea.
			var cmd tea.Cmd
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		}

	case cursor.BlinkMsg:
		// Textarea should also process cursor blinks.
		var cmd tea.Cmd
		m.textarea, cmd = m.textarea.Update(msg)
		return m, cmd

	default:
		return m, nil
	}
}

func (m chatModel) View() string {
	// if m.ascii != nil {
	// 	p := tea.NewProgram(initialSaveModel(m.ascii.art))
	// 	if _, err := p.Run(); err != nil {
	// 		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	// 	}
	// 	return fmt.Sprintln("")
	// } else {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
	// }
}

func (m chatModel) SendMessage(content string) (*openai.ChatCompletionChoice, error) {
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     "gpt-4o-mini",
		MaxTokens: 100,
		Messages: []openai.ChatCompletionMessage{{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		}},
	}
	resp, err := m.aiClient.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return nil, err
	}
	return &resp.Choices[0], nil
}

func storedAsciiArt() tea.Msg {
	return asciiMsg(true)
}
