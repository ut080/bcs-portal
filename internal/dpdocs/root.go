package dpdocs

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/logging"
)

var logLevel string

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dpdocs",
	Short: "Helper tools for Civil Air Patrol admin and org officers",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "")

	logging.InitLogging(logLevel, true)
	config.InitConfig()
}
