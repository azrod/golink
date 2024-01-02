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

// labelCmd represents the label command.
var getLabelCmd = &cobra.Command{
	Use:     "label",
	Aliases: []string{"la"},
	Short:   "Get labels",
	Long:    `Return a list of labels or a single label by name or ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		sdk.SetNamespace(globalFlagNamespace)

		// Create a new context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), globalTimeout())
		defer cancel()

		var labels []struct {
			Label models.Label
			Links []models.Link
		}

		if len(args) > 0 {
			if models.IsValidUUID(args[0]) {
				// Arg is not an ID, try to find the label by name
				label, err := sdk.GetLabelByID(ctx, models.LabelID(args[0]))
				if err != nil {
					log.Default().Printf("Failed to get label: %s", err)
					return
				}

				labels = append(labels, struct {
					Label models.Label
					Links []models.Link
				}{Label: label})
			} else {
				label, err := sdk.GetLabelByName(ctx, args[0])
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
			listLabels, err := sdk.GetLabels(ctx)
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
			links, err := sdk.GetLinksAssociatedToLabel(ctx, label.Label.ID)
			if err != nil {
				log.Default().Printf("Failed to get links associated with label: %s", err)
				return
			}
			labels[i].Links = links
		}

		switch globalFlagOutput {
		case globalFlagOutputShort:
			w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
			fs := "%s\t%s\n"
			fmt.Fprintf(w, fs, "NAME", "LINKS")

			for _, l := range labels {
				fmt.Fprintf(w, fs, l.Label.Name, fmt.Sprintf("%d", len(l.Links)))
			}

			w.Flush()

		case globalFlagOutputWide:
			w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
			fs := "%s\t%s\t%s\t%s\n"
			fmt.Fprintf(w, fs, "ID", "NAME", "LINKS", "COLOR")

			for _, l := range labels {
				fmt.Fprintf(w, fs, l.Label.ID, l.Label.Name, fmt.Sprintf("%d", len(l.Links)), l.Label.Color)
			}

			w.Flush()
		}
	},
}

func init() {
	getCmd.AddCommand(getLabelCmd)
}
