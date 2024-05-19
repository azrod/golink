package glctl

import (
	"context"
	"log"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	golink "github.com/azrod/golink/sdk"
)

// version that can be overwritten by a release process.
var version = "dev"

// commit that can be overwritten by a release process.
var commit = "none"

// date that can be overwritten by a release process.
var date = "unknown"

var apiURL = "http://localhost:8081"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "glctl",
	Short: "glctl is a CLI for golink",
	Long:  `glctl is a CLI for golink. It allows you to manage golink from the command line.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Create a new context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), globalTimeout())
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		cancel()
		os.Exit(1)
	}

	cancel()
}

var (
	globalFlagOutput, globalFlagNamespace string
	globalFlagDebug                       bool
	globalFlagTimeout                     int
	globalTimeout                         = func() time.Duration {
		return time.Duration(globalFlagTimeout) * time.Second
	}
)

const (
	globalFlagOutputShort = "short"
	globalFlagOutputWide  = "wide"

	dirName = ".golink"
)

var (
	cfgHost string
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)
	log.Default().SetFlags(log.Default().Flags() &^ (log.Ldate | log.Ltime))

	rootCmd.PersistentFlags().StringVarP(&globalFlagOutput, "output", "o", globalFlagOutputShort, "output format")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golink/config.yaml)")

	// * Config
	rootCmd.PersistentFlags().BoolVar(&globalFlagDebug, "debug", false, "debug mode")
	if err := viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug")); err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().IntVar(&globalFlagTimeout, "timeout", 10, "timeout in seconds")
	if err := viper.BindPFlag("timeout", rootCmd.PersistentFlags().Lookup("timeout")); err != nil {
		log.Fatal(err)
	}

	// * Namespace
	rootCmd.PersistentFlags().StringVarP(&globalFlagNamespace, "namespace", "n", "default", "namespace")
	if err := rootCmd.RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
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
	}); err != nil {
		log.Fatal(err)
	}

	if err := viper.BindPFlag("namespace", rootCmd.PersistentFlags().Lookup("namespace")); err != nil {
		log.Fatal(err)
	}

	// * Host
	rootCmd.PersistentFlags().StringVar(&cfgHost, "host", apiURL, "golink host")
	if err := viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host")); err != nil {
		log.Fatal(err)
	}
	viper.SetDefault("host", apiURL)

	rootCmd.AddGroup(
		&cobra.Group{
			ID:    "cmd",
			Title: "GoLink Commands",
		},
		&cobra.Group{
			ID:    "other",
			Title: "Other Commands",
		},
	)
}

// initConfig.
func initConfig() {
	viper.SetConfigType("yaml")
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("Unable to find home directory: %s", err)
		}

		// Search config in $HOME/.golink
		viper.AddConfigPath(home + "/" + dirName)
		viper.SetConfigName("config")

		// Check if directory exists.
		if _, err := os.Stat(home + "/" + dirName); os.IsNotExist(err) {
			if err := os.MkdirAll(home+"/"+dirName, 0o755); err != nil {
				log.Printf("Can't create config directory(dir:%s): %s", home+"/"+dirName, err)
				os.Exit(1)
			}
		}

		// Check if config file exists.
		if _, err := os.Stat(home + "/.golink/config.yaml"); os.IsNotExist(err) {
			if err := viper.SafeWriteConfig(); err != nil {
				log.Printf("Can't create config: %s", err)
				os.Exit(1)
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Can't read config: %s", err)
		os.Exit(1)
	}
}

func initSDK() *golink.Client {
	return golink.New(viper.Get("host").(string), viper.Get("debug").(bool), viper.Get("namespace").(string))
}
