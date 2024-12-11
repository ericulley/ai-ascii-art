/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	db "github.com/ericulley/ascii/data"
	"github.com/spf13/cobra"
)


var from string
var to string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update-name",
	Short: "This commands updates the name of an ascii art record",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if from == "" || to == "" {
			fmt.Println("Please specify the name of the ascii art to update [--from] and the new name [--to]")
		} else if from == to {
			fmt.Println("Cannot specify the same name for the ascii art to update [--from] and the new name [--to]")
		} else {
			db.UpdateArt(from, to)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&from, "from", "f", "", "Specify the name of the ascii art to update")
	updateCmd.Flags().StringVarP(&to, "to", "t", "", "Specify the new name to update the ascii art to")
}
