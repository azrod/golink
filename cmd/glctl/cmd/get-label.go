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

// labelCmd represents the label command.
var getLabelCmd = &cobra.Command{
	Use:     "label",
	Aliases: []string{"la"},
	Short:   "Get labels",
	Long:    `Return a list of labels or a single label by name or ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		sdk := initSDK()
		sdk.SetNamespace(globalFlagNamespace)

		var labels []struct {
			Label models.Label
			Links []models.Link
		}

		if len(args) > 0 {
			if models.IsValidUUID(args[0]) {
				// Arg is not an ID, try to find the label by name
				label, err := sdk.GetLabelByID(cmd.Context(), models.LabelID(args[0]))
				if err != nil {
					log.Default().Printf("Failed to get label: %s", err)
					return
				}

				labels = append(labels, struct {
					Label models.Label
					Links []models.Link
				}{Label: label})
			} else {
				label, err := sdk.GetLabelByName(cmd.Context(), args[0])
				if err != nil {
					log.Default().Printf("Failed to get label: %s", err)
					return
				}

				labels = append(labels, struct {
					Label models.Label
					Links []models.Link
				}{Label: label})
			}
		} else {
			listLabels, err := sdk.GetLabels(cmd.Context())
			if err != nil {
				log.Default().Printf("Failed to get labels: %s", err)
				return
			}
			for _, label := range listLabels {
				labels = append(labels, struct {
					Label models.Label
					Links []models.Link
				}{Label: label})
			}
		}

		for i, label := range labels {
			links, err := sdk.GetLinksAssociatedToLabel(cmd.Context(), label.Label.ID)
			if err != nil {
				log.Default().Printf("Failed to get links associated with label: %s", err)
				return
			}
			labels[i].Links = links
		}

		p := print.New()
		defer p.PrintTable()

		switch globalFlagOutput {
		case globalFlagOutputShort:
			p.SetHeader("NAME", "LINKS")

			for _, l := range labels {
				p.AddFields(l.Label.Name, len(l.Links))
			}

		case globalFlagOutputWide:
			p.SetHeader("ID", "NAME", "LINKS", "COLOR")

			for _, l := range labels {
				p.AddFields(l.Label.ID, l.Label.Name, len(l.Links), l.Label.Color)
			}
		}
	},
}

func init() {
	getCmd.AddCommand(getLabelCmd)
}
