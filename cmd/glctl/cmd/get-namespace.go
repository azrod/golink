/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/azrod/golink/models"
)

// getNamespaceCmd represents the namespace command.
var getNamespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"ns"},
	Short:   "Get namespaces",
	Long:    `List all namespaces or get a specific namespace by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		sdk.SetNamespace(globalFlagNamespace)

		// Create a new context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), globalTimeout())
		defer cancel()

		var err error
		nss := []models.Namespace{}

		if len(args) > 0 {
			ns, err := sdk.GetNamespace(ctx, args[0])
			if err != nil {
				log.Default().Printf("Failed to get namespace: %s", err)
				return
			}

			nss = append(nss, ns)
		} else {
			nss, err = sdk.GetNamespaces(ctx)
			if err != nil {
				log.Default().Printf("Failed to get namespaces: %s", err)
				return
			}
		}

		switch globalFlagOutput {
		case globalFlagOutputShort, globalFlagOutputWide:
			w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
			fs := "%s\t%s\t%s\n"
			fmt.Fprintf(w, fs, "NAME", "STATUS", "LINKS")

			for _, l := range nss {
				fmt.Fprintf(w, fs, l.Name, l.Enabled, fmt.Sprintf("%d", len(l.Links)))
			}

			w.Flush()
		}
	},
}

func init() {
	getCmd.AddCommand(getNamespaceCmd)
}
