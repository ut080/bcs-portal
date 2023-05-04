package dpdocs

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/attendance"
	"github.com/ut080/bcs-portal/internal/logging"
)

var attMbrReport string
var attOutfile string

var attendanceCmd = &cobra.Command{
	Use:   "attendance [TableOfOrgFile] [LogDate]",
	Short: "Generate a barcode attendance log from CAPWATCH and unit TO data.",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		logDateStr := args[1]
		logDate, err := time.Parse("2006-01-02", logDateStr)
		if err != nil {
			logging.Error().Err(err).Msg("Invalid date format for log date, use ISO 8601")
			os.Exit(1)
		}

		if attOutfile == "" {
			attOutfile = fmt.Sprintf("%s.pdf", logDateStr)
		}

		err = attendance.BuildBarcodeLog(args[0], attOutfile, attMbrReport, logDate)
		if err != nil {
			logging.Error().Err(err).Msg("Failed to generate barcode attendance log")
			os.Exit(1)
		}
	},
}

func init() {
	attendanceCmd.Flags().StringVarP(&attOutfile, "out", "o", "", "output file path (defaults to the log date)")
	attendanceCmd.Flags().StringVarP(&attMbrReport, "membership-report", "r", "", "file path to eServices Membership report (skips CAPWATCH access)")

	rootCmd.AddCommand(attendanceCmd)
}
