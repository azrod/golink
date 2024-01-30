/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package glctl

import (
	"log"

	"github.com/orange-cloudavenue/common-go/print"
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

		p := print.New()
		defer p.PrintTable()

		switch globalFlagOutput {
		case globalFlagOutputShort:
			p.SetHeader("NAMESPACE", "NAME", "PATH", "TARGET URL", "STATUS")

			for _, l := range links {
				if len(l.TargetURL) > 50 {
					l.TargetURL = l.TargetURL[:50] + "..."
				}
				p.AddFields(globalFlagNamespace, l.Name, l.SourcePath, l.TargetURL, l.Enabled)
			}

		case globalFlagOutputWide:
			p.SetHeader("NAMESPACE", "NAME", "PATH", "TARGET URL", "STATUS", "LABELS")

			for _, l := range links {
				p.AddFields(globalFlagNamespace, l.Name, l.SourcePath, l.TargetURL, l.Enabled, l.Labels)
			}
		}
	},
}

func init() {
	getCmd.AddCommand(getLinkCmd)
}
