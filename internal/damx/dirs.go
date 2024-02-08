package damx

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/logging"
)

var dirsCmd = &cobra.Command{
	Use:   "dirs [outputDir]",
	Short: "Generate a directory hierarchy for squadron personnel (currently disabled)",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run:   runDirsCmd,
}

func runDirsCmd(cmd *cobra.Command, args []string) {
	// TODO: Integrate this with file plan/local database
	logging.Error().Msg("dirs command is currently disabled")
	os.Exit(1)

	/*
		logger := logging.Logger{}

		var err error

		// TODO: Check that all three report flags are set if any are set

		err = personnel.CreateDirectories(mbrReports, args[1])
		if err != nil {
			logging.Error().Err(err).Msg("Failed to create personnel file directories")
			os.Exit(1)
		}
	*/
}

func init() {
	RootCmd.AddCommand(dirsCmd)
}
