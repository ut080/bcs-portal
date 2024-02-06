package damx

import (
	"os"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/internal/personnel"
	"github.com/ut080/bcs-portal/pkg/org"
)

var cdtMbrReportStr string
var csMbrReportStr string
var sMbrReportStr string

var dirsCmd = &cobra.Command{
	Use:   "dirs [outputDir]",
	Short: "Generate a directory hierarchy for squadron personnel",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
		if cdtMbrReportStr != "" {
			smReport, err = files.NewFile(sMbrReportStr, logger.DefaultLogger())
			if err != nil {
				logging.Error().Err(err).Str("sMbrReportStr", sMbrReportStr).Msg("Failed to create file reference for Senior Member report")
				os.Exit(1)
			}
		}

		// TODO: Check that all three report flags are set if any are set
		mbrReports := map[org.MemberType]files.File{
			org.CadetMember:        cdtReport,
			org.CadetSponsorMember: csmReport,
			org.SeniorMember:       smReport,
		}

		err = personnel.CreateDirectories(mbrReports, args[1])
		if err != nil {
			logging.Error().Err(err).Msg("Failed to create personnel file directories")
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(dirsCmd)

	dirsCmd.Flags().StringVarP(&cdtMbrReportStr, "cdt-report", "c", "", "file path to eServices Cadet Membership report (skips CAPWATCH access). MUST be used with --csm-report and --sm-report.")
	dirsCmd.Flags().StringVarP(&csMbrReportStr, "csm-report", "s", "", "file path to eServices Cadet Sponsor Membership report (skips CAPWATCH access). MUST be used with --cdt-report and --sm-report.")
	dirsCmd.Flags().StringVarP(&cdtMbrReportStr, "sm-report", "r", "", "file path to eServices Cadet Membership report (skips CAPWATCH access). MUST be used with --cdt-report and --csm-report.")
}
