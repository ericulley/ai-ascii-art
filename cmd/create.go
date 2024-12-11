/*
Copyright © 2024 Eric Culley <https://github.com/ericulley>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ericulley/ascii/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Opens a chat session with AI to generate an ascii art",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.NewChatModel())
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
