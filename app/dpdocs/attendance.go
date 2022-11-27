package dpdocs

import (
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/app/attendance"
	"github.com/ut080/bcs-portal/app/logging"
)

var capwatchPassword string // TODO: Handle this without having to type password in the clear on the command line
var outfile string

var attendanceCmd = &cobra.Command{
	Use:   "attendance [TableOfOrgFile] [LogDate]",
	Short: "Generate a barcode attendance log from CAPWATCH and unit TO data.",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		logDate, err := time.Parse("2006-01-02", args[1])
		if err != nil {
			logging.Error().Err(err).Msg("Invalid date format for log date, use ISO 8601")
			os.Exit(1)
		}

		err = attendance.BuildBarcodeLog(args[0], outfile, capwatchPassword, logDate)
		if err != nil {
			logging.Error().Err(err).Msg("Failed to generate barcode attendance log")
			os.Exit(1)
		}
	},
}

func init() {
	attendanceCmd.Flags().StringVarP(&capwatchPassword, "password", "p", "", "eServices password")
	attendanceCmd.Flags().StringVarP(&outfile, "out", "o", "attendance.pdf", "output file path")

	rootCmd.AddCommand(attendanceCmd)
}
