package dpdocs

import (
	"fmt"
	"os"
	"time"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/attendance"
	"github.com/ut080/bcs-portal/internal/logging"
)

var attMbrReportStr string
var attOutfileStr string

var attendanceCmd = &cobra.Command{
	Use:   "attendance [TableOfOrgFile] [LogDate]",
	Short: "Generate a barcode attendance log from CAPWATCH and unit TO data.",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.Logger{}

		toCfg, err := files.NewFile(args[0], logger.DefaultLogger())
		if err != nil {
			logging.Error().Err(err).Str("toCfg", args[0]).Msg("Failed to create file reference for TO cfg")
			os.Exit(1)
		}

		logDateStr := args[1]
		logDate, err := time.Parse("2006-01-02", logDateStr)
		if err != nil {
			logging.Error().Err(err).Msg("Invalid date format for log date, use ISO 8601")
			os.Exit(1)
		}

		var attOutfile files.File
		if attOutfileStr == "" {
			attOutfileStr = fmt.Sprintf("%s.pdf", logDateStr)
		}
		attOutfile, err = files.NewFile(attOutfileStr, logger.DefaultLogger())
		if err != nil {
			logging.Error().Err(err).Str("attOutfileStr", attOutfileStr).Msg("Failed to create file reference for log")
			os.Exit(1)
		}

		var attMbrReport files.File
		if attMbrReportStr != "" {
			attMbrReport, err = files.NewFile(attMbrReportStr, logger.DefaultLogger())
			if err != nil {
				logging.Error().Err(err).Str("attMbrReportStr", attMbrReportStr).Msg("Failed to create file reference for Member report")
				os.Exit(1)
			}
		}

		err = attendance.BuildBarcodeLog(toCfg, attOutfile, attMbrReport, logDate, logger)
		if err != nil {
			logging.Error().Err(err).Msg("Failed to generate barcode attendance log")
			os.Exit(1)
		}
	},
}

func init() {
	attendanceCmd.Flags().StringVarP(&attOutfileStr, "out", "o", "", "output file path (defaults to the log date)")
	attendanceCmd.Flags().StringVarP(&attMbrReportStr, "membership-report", "r", "", "file path to eServices Membership report (skips CAPWATCH access)")

	rootCmd.AddCommand(attendanceCmd)
}
