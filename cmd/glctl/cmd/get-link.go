/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package glctl

import (
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
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		sdk := initSDK()
		links, err := sdk.GetLinks(cmd.Context())
		if err != nil {
			log.Default().Printf("Failed to get links: %s", err)
			return nil, cobra.ShellCompDirectiveError
		}

		var names []string
		for _, l := range links {
			names = append(names, l.Name)
		}

		return names, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			sdk   = initSDK()
			links []models.Link
			err   error
		)

		spin.Start()

		if len(args) > 0 {
			for _, arg := range args {
				var (
					link models.Link
					err  error
				)
				// Arg is not an ID, try to find the link by name
				if models.IsValidUUID(arg) {
					link, err = sdk.GetLinkByID(cmd.Context(), models.LinkID(arg))
				} else {
					link, err = sdk.GetLinkByName(cmd.Context(), arg)
				}
				if err != nil {
					log.Default().Printf("Failed to get link: %s", err)
					return
				}

				links = append(links, link)
			}
		} else {
			links, err = sdk.GetLinks(cmd.Context())
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
				if len(l.TargetURL) > 50 {
					l.TargetURL = l.TargetURL[:50] + "..."
				}
				fmt.Fprintf(w, fs, globalFlagNamespace, l.Name, l.SourcePath, l.TargetURL, l.Enabled.String())
			}
			spin.Stop()
			w.Flush()

		case globalFlagOutputWide:
			w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
			fs := "%s\t%s\t%s\t%s\t%s\t%s\n"
			fmt.Fprintf(w, fs, "NAMESPACE", "NAME", "PATH", "TARGET URL", "STATUS", "LABELS")

			for _, l := range links {
				fmt.Fprintf(w, fs, globalFlagNamespace, l.Name, l.SourcePath, l.TargetURL, l.Enabled.String(), l.Labels)
			}
			spin.Stop()
			w.Flush()
		}
	},
}

func init() {
	getCmd.AddCommand(getLinkCmd)
}
