package dpdocs

import (
	"fmt"
	"os"
	"time"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/attendance"
	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/org"
)

var attCdtMbrReportStr string
var attCSMbrReportStr string
var attSMbrReportStr string
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

		var cdtReport files.File
		if attCdtMbrReportStr != "" {
			cdtReport, err = files.NewFile(attCdtMbrReportStr, logger.DefaultLogger())
			if err != nil {
				logging.Error().Err(err).Str("attMbrReportStr", attCdtMbrReportStr).Msg("Failed to create file reference for Cadet Member report")
				os.Exit(1)
			}
		}

		var csmReport files.File
		if attCSMbrReportStr != "" {
			csmReport, err = files.NewFile(attCSMbrReportStr, logger.DefaultLogger())
			if err != nil {
				logging.Error().Err(err).Str("attCSMbrReportStr", attCSMbrReportStr).Msg("Failed to create file reference for Cadet Sponsor Member report")
				os.Exit(1)
			}
		}

		var smReport files.File
		if attSMbrReportStr != "" {
			smReport, err = files.NewFile(attSMbrReportStr, logger.DefaultLogger())
			if err != nil {
				logging.Error().Err(err).Str("attSMbrReportStr", attSMbrReportStr).Msg("Failed to create file reference for Senior Member report")
				os.Exit(1)
			}
		}

		// TODO: Check that all three report flags are set if any are set
		mbrReports := map[org.MemberType]files.File{
			org.CadetMember:        cdtReport,
			org.CadetSponsorMember: csmReport,
			org.SeniorMember:       smReport,
		}

		err = attendance.BuildBarcodeLog(toCfg, attOutfile, mbrReports, logDate, logger)
		if err != nil {
			logging.Error().Err(err).Msg("Failed to generate barcode attendance log")
			os.Exit(1)
		}
	},
}

func init() {
	attendanceCmd.Flags().StringVarP(&attOutfileStr, "out", "o", "", "output file path (defaults to the log date)")
	attendanceCmd.Flags().StringVarP(&attCdtMbrReportStr, "cdt-report", "c", "", "file path to eServices Cadet Membership report (skips CAPWATCH access). MUST be used with --csm-report and --sm-report.")
	attendanceCmd.Flags().StringVarP(&attCSMbrReportStr, "csm-report", "s", "", "file path to eServices Cadet Sponsor Membership report (skips CAPWATCH access). MUST be used with --cdt-report and --sm-report.")
	attendanceCmd.Flags().StringVarP(&attSMbrReportStr, "sm-report", "r", "", "file path to eServices Cadet Membership report (skips CAPWATCH access). MUST be used with --cdt-report and --csm-report.")

	rootCmd.AddCommand(attendanceCmd)
}
