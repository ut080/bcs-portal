package dpdocs

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/logging"
)

var logOutFileStr string

var fileplanCmd = &cobra.Command{
	Use:   "attendance [date]",
	Short: "Generate a Barcode Attendance Log for the given meeting date.",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run:   runAttendance,
}

func init() {
	fileplanCmd.Flags().StringVarP(&logOutFileStr, "out", "o", "", "output file path for attendance log (defaults to <SquadronName>-<MeetingDate>.pdf).")

	rootCmd.AddCommand(fileplanCmd)
}

func runAttendance(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

	logger.Error().Msg("Command not implemented.")
	os.Exit(1)

}
