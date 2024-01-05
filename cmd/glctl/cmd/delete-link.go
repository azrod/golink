package glctl

import (
	"log"

	"github.com/go-resty/resty/v2"
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
		sdk.SetNamespace(globalFlagNamespace)

		id := args[0]

		if !models.IsValidUUID(id) {
			// Arg is not an ID, try to find the link by name
			rFromName, err := resty.New().
				SetBaseURL(apiURL).
				R().
				SetDebug(globalFlagDebug).
				SetResult(&models.Link{}).
				Get("/api/v1/link/name/" + id)
			if err != nil {
				log.Default().Printf("Failed to get link: %s", err)
				return
			}

			if rFromName.IsError() {
				log.Default().Printf("Failed to get link: %s", rFromName.Error())
				return
			}

			id = rFromName.Result().(*models.Link).ID.String()
		}

		v := &models.Link{
			LinkRequest: models.LinkRequest{
				Name:       args[0],
				SourcePath: args[1],
				TargetURL:  args[2],
				Enabled:    models.Enabled(!addCmdFlagDisable),
			},
			// TODO add labels and group
		}

		r, err := resty.New().
			SetBaseURL(apiURL).
			R().
			SetDebug(globalFlagDebug).
			SetResult(&models.Link{}).
			SetBody(v).
			SetPathParam("id", id).
			Delete("/api/v1/link/{id}")
		if err != nil {
			log.Default().Printf("Failed to deletelink: %s", err)
			return
		}

		if r.IsError() {
			log.Default().Printf("Failed to deletelink: %s", r.Error())
			return
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
