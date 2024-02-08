package damx

import (
	"os"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/internal/update"
	"github.com/ut080/bcs-portal/pkg/org"
)

var cdtMbrReportStr string
var csMbrReportStr string
var sMbrReportStr string

var updateCmd = &cobra.Command{
	Use:   "update --cdt-report=[Cadet Membership Report] --csm-report=[CSM Membership Report] --sm-report[SM Membership Report]",
	Short: "Generate a Barcode Attendance Log for the given meeting date.",
	Long:  ``,
	Args:  cobra.ExactArgs(0),
	Run:   runLocalUpdateCmd,
}

func runLocalUpdateCmd(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

	var err error
	var cdtReport files.File
	if cdtMbrReportStr != "" {
		cdtReport, err = files.NewFile(cdtMbrReportStr, logger.DefaultLogger())
		if err != nil {
			logging.Error().Err(err).Str("attMbrReportStr", cdtMbrReportStr).Msg("Failed to create file reference for Cadet Member report")
			os.Exit(1)
		}
	}

	var csmReport files.File
	if csMbrReportStr != "" {
		csmReport, err = files.NewFile(csMbrReportStr, logger.DefaultLogger())
		if err != nil {
			logging.Error().Err(err).Str("csMbrReportStr", csMbrReportStr).Msg("Failed to create file reference for Cadet Sponsor Member report")
			os.Exit(1)
		}
	}

	var smReport files.File
	if sMbrReportStr != "" {
		smReport, err = files.NewFile(sMbrReportStr, logger.DefaultLogger())
		if err != nil {
			logging.Error().Err(err).Str("sMbrReportStr", sMbrReportStr).Msg("Failed to create file reference for Senior Member report")
			os.Exit(1)
		}
	}

	mbrReports := map[org.MemberType]files.File{
		org.CadetMember:        cdtReport,
		org.CadetSponsorMember: csmReport,
		org.SeniorMember:       smReport,
	}

	err = update.Update(false, mbrReports)
	if err != nil {
		logging.Error().Err(err).Msg("Error occurred while trying to update local database.")
	}

	logger.Error().Msg("Command not implemented.")
	os.Exit(1)
}

func init() {
	updateCmd.Flags().StringVarP(&cdtMbrReportStr, "cdt-report", "c", "", "file path to eServices Cadet Membership report (ignores local database). MUST be used with --csm-report and --sm-report.")
	updateCmd.Flags().StringVarP(&csMbrReportStr, "csm-report", "s", "", "file path to eServices Cadet Sponsor Membership report (ignores local database). MUST be used with --cdt-report and --sm-report.")
	updateCmd.Flags().StringVarP(&sMbrReportStr, "sm-report", "r", "", "file path to eServices Cadet Membership report (ignores local database). MUST be used with --cdt-report and --csm-report.")

	RootCmd.AddCommand(updateCmd)
}
