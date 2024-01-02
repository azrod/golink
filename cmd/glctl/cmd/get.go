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
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
