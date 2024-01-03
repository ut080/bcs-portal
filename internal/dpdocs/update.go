package dpdocs

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/logging"
)

var localUpdateCmd = &cobra.Command{
	Use:   "local-update [Cadet Membership Report] [CSM Membership Report] [SM Membership Report]",
	Short: "Generate a Barcode Attendance Log for the given meeting date.",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	Run:   runLocalUpdate,
}

func init() {
	rootCmd.AddCommand(localUpdateCmd)
}

func runLocalUpdate(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

	logger.Error().Msg("Command not implemented.")
	os.Exit(1)
}
