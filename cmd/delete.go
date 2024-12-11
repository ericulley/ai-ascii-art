/*
Copyright © 2024 Eric Culley <https://github.com/ericulley>
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
	Short: "Deletes an ascii art record",
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
