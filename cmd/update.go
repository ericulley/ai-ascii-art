/*
Copyright © 2024 Eric Culley <https://github.com/ericulley>
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
	Short: "Updates the name of an ascii art record",
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
