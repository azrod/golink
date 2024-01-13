/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package glctl

import (
	"log"

	"github.com/spf13/cobra"
)

// namespaceCmd represents the namespace command.
var addNamespaceCmd = &cobra.Command{
	Use:     "namespace [NAME]",
	Aliases: []string{"ns"},
	Short:   "Add namespaces",
	Long:    `Add namespaces are used to add namespaces to the database.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		return cobra.MinimumNArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		sdk := initSDK()

		_, err := sdk.CreateNamespace(cmd.Context(), args[0])
		if err != nil {
			log.Default().Printf("Failed to add namespace: %s", err)
			return
		}

		log.Default().Printf("Successfully added namespace %s !", args[0])
	},
}

func init() {
	addCmd.AddCommand(addNamespaceCmd)
}
