package dpdocs

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/logging"
)

var logLevel string
var outfile string
var mbrReport string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dpdocs",
	Short: "Helper tools for Civil Air Patrol admin and personnel officers",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "")
	rootCmd.PersistentFlags().StringVarP(&outfile, "out", "o", "", "output file path (defaults to the log date)")
	rootCmd.PersistentFlags().StringVarP(&mbrReport, "membership-report", "r", "", "file path to eServices Membership report (skips CAPWATCH access)")

	logging.InitLogging(logLevel, true)
	config.InitConfig()
}
