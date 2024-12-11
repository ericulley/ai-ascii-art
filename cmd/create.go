/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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
	Short: "This command opens a chat session with AI to generate an ascii art",
	Long:  `This commands sends a message to chatGPT`,
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
