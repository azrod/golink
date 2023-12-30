/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"

	golink "github.com/azrod/golink/sdk"
)

var sdk *golink.Client

var apiURL = "http://localhost:8081"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
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
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&globalFlagOutput, "output", "o", globalFlagOutputShort, "output format")
	rootCmd.PersistentFlags().BoolVar(&globalFlagDebug, "debug", false, "debug mode")
	rootCmd.PersistentFlags().IntVar(&globalFlagTimeout, "timeout", 10, "timeout in seconds")
	rootCmd.PersistentFlags().StringVarP(&globalFlagNamespace, "namespace", "n", "", "namespace")

	if globalFlagNamespace == "" {
		globalFlagNamespace = "default"
	}

	log.Default().SetFlags(log.Default().Flags() &^ (log.Ldate | log.Ltime))

	sdk = golink.New(apiURL, &globalFlagDebug, &globalFlagNamespace)
}
