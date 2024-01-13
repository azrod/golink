package glctl

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/azrod/golink/models"
)

var addLabelCmdFlagColor string

var addLabelCmdListColors = func() string {
	var colors string
	for _, c := range models.Colors {
		colors += fmt.Sprintf("\t- %s\n", c)
	}
	return colors
}

// labelCmd represents the label command.
var addLabelCmd = &cobra.Command{
	Use:     "label",
	Aliases: []string{"la"},
	Short:   "Add a label [NAME]",
	Long: `Add a label to the database.
	
	Basic usage:
	$ golink add label [NAME]
	
	Add a label with a color:
	$ golink add label [NAME] --color [COLOR]
	
	Colors can be one of the following:
	` + addLabelCmdListColors(),

	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		return cobra.MinimumNArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		sdk := initSDK()

		v := models.LabelRequest{
			Name: args[0],
		}

		if addLabelCmdFlagColor != "" {
			if ok, err := models.IsValidColor(models.ColorName(addLabelCmdFlagColor)); !ok || err != nil {
				log.Default().Printf("Invalid color %s\n", addLabelCmdFlagColor)
				return
			}
			v.Color = models.ColorName(addLabelCmdFlagColor)
		}

		if _, err := sdk.AddLabel(cmd.Context(), v); err != nil {
			log.Default().Printf("Failed to add label: %s\n", err)
			return
		}

		log.Default().Printf("Label %s added\n", args[0])
	},
}

func init() {
	addCmd.AddCommand(addLabelCmd)

	// Add flag color
	addLabelCmd.Flags().StringVarP(&addLabelCmdFlagColor, "color", "c", "", "Color of the label")
}
