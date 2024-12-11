/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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
	Short: "lists the saved ascii art",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
