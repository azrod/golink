/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package glctl

import (
	"log"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command.
var versionCmd = &cobra.Command{
	Use:     "version",
	GroupID: "other",
	Short:   "Returns the version of the application",
	Long:    `Returns the version, commit hash and build date of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Version: %s\nCommit: %s\nBuild Date: %s\n", version, commit, date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
