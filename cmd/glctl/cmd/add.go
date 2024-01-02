/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package glctl

import (
	"github.com/spf13/cobra"
)

var addCmdFlagDisable bool

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add commands",
	GroupID: "cmd",
	Long:    `Add commands are used to add information to the database.`,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
