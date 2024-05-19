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

// getNamespaceCmd represents the namespace command.
var getNamespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"ns"},
	Short:   "Get namespaces",
	Long:    `List all namespaces or get a specific namespace by name.`,
	Example: `
# List all namespaces
$> glctl get namespace
	  
# Get a specific namespace
$> glctl get namespace [NAME]

# Get a multiple namespaces
$> glctl get namespace [NAME] [NAME] [NAME]`,
	ValidArgsFunction: func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		sdk := initSDK()
		nss, err := sdk.GetNamespaces(cmd.Context())
		if err != nil {
			log.Default().Printf("Failed to get namespaces: %s", err)
			return nil, cobra.ShellCompDirectiveError
		}

		var names []string
		for _, ns := range nss {
			names = append(names, ns.Name)
		}

		return names, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			sdk = initSDK()
			err error
			nss = []models.Namespace{}
		)

		if len(args) > 0 {
			for _, name := range args {
				ns, err := sdk.GetNamespace(cmd.Context(), name)
				if err != nil {
					log.Default().Printf("Failed to get namespace: %s", err)
					return
				}

				nss = append(nss, ns)
			}
		} else {
			nss, err = sdk.GetNamespaces(cmd.Context())
			if err != nil {
				log.Default().Printf("Failed to get namespaces: %s", err)
				return
			}
		}

		switch globalFlagOutput {
		case globalFlagOutputShort, globalFlagOutputWide:
			p := print.New()
			p.SetHeader("NAME", "STATUS", "LINKS")
			for _, l := range nss {
				p.AddFields(l.Name, l.Enabled, len(l.Links))
			}
			p.PrintTable()
		}
	},
}

func init() {
	getCmd.AddCommand(getNamespaceCmd)
}
