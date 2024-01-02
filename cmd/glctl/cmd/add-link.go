/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package glctl

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/azrod/golink/models"
)

// linkCmd represents the link command.
var (
	addLinkCmd = &cobra.Command{
		Use:     "link [NAME] [PATH] [TARGET URL]",
		Aliases: []string{"li"},
		Short:   "Add a link",
		Long:    `Add a link to the database.`,
		// GroupID: "link",
		Args: func(cmd *cobra.Command, args []string) error {
			// Optionally run one of the validators provided by cobra
			return cobra.MinimumNArgs(3)(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			sdk.SetNamespace(globalFlagNamespace)

			// Create a new context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), globalTimeout())
			defer cancel()

			v := models.LinkRequest{
				Name:       args[0],
				SourcePath: args[1],
				TargetURL:  args[2],
				Enabled:    models.Enabled(!addCmdFlagDisable),
			}

			_, err := sdk.CreateLink(ctx, v)
			if err != nil {
				log.Default().Printf("Failed to add link: %s", err)
				return
			}

			log.Default().Printf("Successfully added link %s !", args[0])
		},
	}
)

func init() {
	addCmd.AddCommand(addLinkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().BoolVar(&addCmdFlagDisable, "disable", false, "Set the link as disabled")
}
