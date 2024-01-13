package glctl

import (
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// goCmd represents the go command.
var goCmd = &cobra.Command{
	Use:     "go",
	Short:   "Open a link in the default browser",
	GroupID: "cmd",
	Long:    `Perform a GET request on the link and open the target URL in the default browser.`,
	Example: `
# Open a link in the default browser
$> glctl go [LINKNAME]
$> glctl go [NAMESPACE/LINKNAME]
`,
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.ExactArgs(1)(cmd, args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		sdk := initSDK()
		nss, err := sdk.GetNamespaces(cmd.Context())
		if err != nil {
			log.Default().Printf("Failed to get namespaces: %s", err)
			return nil, cobra.ShellCompDirectiveError
		}

		var names []string

		// name format is namespace/name
		for _, ns := range nss {
			for _, link := range ns.Links {
				names = append(names, ns.Name+"/"+link.Name)
			}
		}

		return names, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		sdk := initSDK()

		var linkName string

		// split args[0] to get namespace and link name
		// format possible are namespace/linkname or linkname (default namespace)
		u := strings.Split(args[0], "/")

		if len(u) == 2 {
			sdk.SetNamespace(u[0])
			linkName = u[1]
		} else {
			sdk.SetNamespace("default")
			linkName = u[0]
		}

		link, err := sdk.GetLinkByName(cmd.Context(), linkName)
		if err != nil {
			log.Default().Printf("Failed to get namespace: %s", err)
			return
		}

		var (
			cmdOS  string
			argsOS []string
		)

		switch runtime.GOOS {
		case "windows":
			cmdOS = "cmd"
			argsOS = []string{"/c", "start"}
		case "darwin":
			cmdOS = "open"
		default: // "linux", "freebsd", "openbsd", "netbsd"
			cmdOS = "xdg-open"
		}
		argsOS = append(argsOS, link.TargetURL)
		if err := exec.Command(cmdOS, argsOS...).Start(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}
