package glctl

import (
	"log"

	"github.com/spf13/cobra"
)

var deleteNamespaceCmdFlagForce bool

// deleteLinkCmd represents the namespace command.
var deleteNamespaceCmd = &cobra.Command{
	Use:     "namespace [NAME]",
	Aliases: []string{"ns"},
	Short:   "Delete a namespace",
	Long:    `Delete a namespace from the database.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		return cobra.ExactArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		sdk := initSDK()

		if _, err := sdk.GetNamespace(cmd.Context(), args[0]); err != nil {
			log.Default().Printf("Failed to get namespace: %s", err)
			return
		}

		if deleteNamespaceCmdFlagForce {
			sdk.SetNamespace(args[0])

			log.Default().Printf("Force deletion of namespace %s detected", args[0])
			links, err := sdk.GetLinks(cmd.Context())
			if err != nil {
				log.Default().Printf("Failed to list links: %s", err)
				return
			}

			if len(links) == 0 {
				log.Default().Printf("No links found in namespace %s", args[0])
			} else {
				log.Default().Printf("Found %d links in namespace %s.\nStarting deletion...", len(links), args[0])
				for _, link := range links {
					if err := sdk.DeleteLink(cmd.Context(), link.ID); err != nil {
						log.Default().Printf("Failed to delete link: %s", err)
						return
					}
					log.Default().Printf("Deleted link %s", link.Name)
				}
			}
		}

		if err := sdk.DeleteNamespace(cmd.Context(), args[0]); err != nil {
			log.Default().Printf("Failed to delete namespace: %s", err)
			return
		}

		log.Default().Printf("Successfully deleted namespace %s !", args[0])
	},
}

func init() {
	deleteNamespaceCmd.Flags().BoolVarP(&deleteNamespaceCmdFlagForce, "force", "f", false, "Force deletion of namespace. This will delete all links in the namespace.")

	deleteCmd.AddCommand(deleteNamespaceCmd)
}
