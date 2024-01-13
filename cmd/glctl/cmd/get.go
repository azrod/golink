/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package glctl

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command.
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get commands",
	GroupID: "cmd",
	Long:    `Get commands are used to retrieve information from the database.`,
	ValidArgs: []string{
		"link",
		"namespace",
		"label",
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
