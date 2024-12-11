/*
Copyright © 2024 Eric Culley <https://github.com/ericulley>
*/
package cmd

import (
	db "github.com/ericulley/ascii/data"

	"github.com/spf13/cobra"
)

var limit int
var names bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows ascii art stored in the database",
	Run: func(cmd *cobra.Command, args []string) {
		if limit != 0 && names {
			db.ListArtNamesWithLimit(limit)
		} else if limit != 0 && !names {
			db.ListArtWithLimit(limit)
		} else if limit == 0 && names {
			db.ListArtNames()
		} else {
			db.ListAllArt()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&limit, "limit", "l", 0, "limit the number of results")
	listCmd.Flags().BoolVarP(&names, "names", "n", false, "list only the name column of the results")
}
