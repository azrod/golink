package glctl

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/azrod/golink/models"
)

// labelCmd represents the label command.
var delLabelCmd = &cobra.Command{
	Use:     "label",
	Aliases: []string{"la"},
	Short:   "Delete a label",
	Long:    `Delete a label from the database.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return cobra.MinimumNArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		sdk.SetNamespace(globalFlagNamespace)

		// Create a new context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), globalTimeout())
		defer cancel()

		if models.IsValidUUID(args[0]) {
			// Arg is not an ID, try to find the label by name
			if err := sdk.DeleteLabelByID(ctx, models.LabelID(args[0])); err != nil {
				log.Default().Printf("Failed to get label: %s", err)
				return
			}
		} else {
			if err := sdk.DeleteLabelByName(ctx, args[0]); err != nil {
				log.Default().Printf("Failed to get label: %s", err)
				return
			}
		}

		fmt.Printf("Label %s deleted\n", args[0]) //nolint: forbidigo
	},
}

func init() {
	deleteCmd.AddCommand(delLabelCmd)
}
