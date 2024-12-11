/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	db "github.com/ericulley/ascii/data"
	"github.com/spf13/cobra"
)

var id int
var name string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if id == 0 && name == "" {
			fmt.Println("Please specify an id [--id] or name [--name] of the ascii art to delete")
		} else if id != 0 && name != "" {
			fmt.Println("Cannot specify both id [--id] and name [--name] of the ascii art to delete")
		} else if id != 0 && name == "" {
			db.DeleteArtById(id)
		} else if id == 0 && name != "" {
			db.DeleteArtByName(name)
		} 
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVar(&id, "id", 0, "Specify the id of the ascii art to delete")
	deleteCmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the ascii art to delete")
}
