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
		sdk := initSDK()

		// Ignore errors, we don't care if the server is not reachable
		vServer, _ := sdk.GetVersion(cmd.Context())

		log.Printf(`Client informations:
  Version: %s
  Commit: %s
  Build Date: %s`,
			version, commit, date)

		if vServer != "" {
			log.Printf(`
Server informations:
  Version: %s
`, vServer)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
