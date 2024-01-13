package glctl

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/azrod/golink/models"
)

// linkCmd represents the link command.
var deleteLinkCmd = &cobra.Command{
	Use:     "link [NAME] | [ID]",
	Aliases: []string{"li"},
	Short:   "Delete a link",
	Long:    `Delete a link from the database.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		return cobra.ExactArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		sdk := initSDK()

		if !models.IsValidUUID(args[0]) {
			if err := sdk.DeleteLinkByName(cmd.Context(), args[0]); err != nil {
				log.Default().Printf("Failed to delete link: %s", err)
				return
			}
		} else {
			if err := sdk.DeleteLink(cmd.Context(), models.LinkID(args[0])); err != nil {
				log.Default().Printf("Failed to delete link: %s", err)
				return
			}
		}

		log.Default().Printf("Successfully deleted link %s !", args[0])
	},
}

func init() {
	deleteCmd.AddCommand(deleteLinkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
