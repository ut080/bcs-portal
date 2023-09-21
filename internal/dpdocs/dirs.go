package dpdocs

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/internal/personnel"
)

var dirsCmd = &cobra.Command{
	Use:   "dirs [eServicesReport] [outputDir]",
	Short: "Generate a directory hierarchy for squadron personnel",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := personnel.CreateDirectories(args[0], args[1])
		if err != nil {
			logging.Error().Err(err).Msg("Failed to create personnel file directories")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(dirsCmd)
}
