/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package glctl

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/azrod/golink/models"
)

// linkCmd represents the link command.
var getLinkCmd = &cobra.Command{
	Use:     "link",
	Aliases: []string{"li"},
	Short:   "Get links",
	Long:    `Return a list of links or a single link by name or ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		sdk.SetNamespace(globalFlagNamespace)

		// Create a new context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), globalTimeout())
		defer cancel()

		var (
			links []models.Link
			err   error
		)

		if len(args) > 0 {
			if models.IsValidUUID(args[0]) {
				// Arg is not an ID, try to find the link by name
				link, err := sdk.GetLinkByID(ctx, models.LinkID(args[0]))
				if err != nil {
					log.Default().Printf("Failed to get link: %s", err)
					return
				}

				links = append(links, link)
			} else {
				link, err := sdk.GetLinkByName(ctx, args[0])
				if err != nil {
					log.Default().Printf("Failed to get link: %s", err)
					return
				}

				links = append(links, link)
			}
		} else {
			links, err = sdk.GetLinks(ctx)
			if err != nil {
				log.Default().Printf("Failed to get links: %s", err)
				return
			}
		}

		switch globalFlagOutput {
		case globalFlagOutputShort:
			w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
			fs := "%s\t%s\t%s\t%s\t%s\n"
			fmt.Fprintf(w, fs, "NAMESPACE", "NAME", "PATH", "TARGET URL", "STATUS")

			for _, l := range links {
				fmt.Fprintf(w, fs, globalFlagNamespace, l.Name, l.SourcePath, l.TargetURL, l.Enabled.String())
			}

			w.Flush()

		case globalFlagOutputWide:
			w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
			fs := "%s\t%s\t%s\t%s\t%s\t%s\n"
			fmt.Fprintf(w, fs, "NAMESPACE", "NAME", "PATH", "TARGET URL", "STATUS", "LABELS")

			for _, l := range links {
				fmt.Fprintf(w, fs, globalFlagNamespace, l.Name, l.SourcePath, l.TargetURL, l.Enabled.String(), l.Labels)
			}

			w.Flush()
		}
	},
}

func init() {
	getCmd.AddCommand(getLinkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
