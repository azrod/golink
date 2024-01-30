package glctl

import (
	"github.com/orange-cloudavenue/common-go/print"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command.
var versionCmd = &cobra.Command{
	Use:     "version",
	GroupID: "other",
	Short:   "Returns the version of the application",
	Long:    `Returns the version, commit hash and build date of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sdk := initSDK()

		// Ignore errors, we don't care if the server is not reachable
		vServer, _ := sdk.GetVersion(cmd.Context())

		p := print.New()
		defer p.PrintTable()
		p.SetHeader("Type", "Version", "Commit", "Build Date")
		p.AddFields("Client", version, commit, date)

		if vServer != "" {
			p.AddFields("Server", vServer, "", "")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
