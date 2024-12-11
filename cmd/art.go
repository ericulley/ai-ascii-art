/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	db "github.com/ericulley/ascii/data"
	"github.com/spf13/cobra"
)

// artCmd represents the art command
var artCmd = &cobra.Command{
	Use:   "art",
	Short: "Return a random ascii art stored in the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		db.Art()
	},
}

func init() {
	rootCmd.AddCommand(artCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// artCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// artCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
