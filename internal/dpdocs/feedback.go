package dpdocs

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/feedback"
	"github.com/ut080/bcs-portal/internal/logging"
)

var fbkMbrReport string
var fbkOutfile string

var feedbackCmd = &cobra.Command{
	Use:   "feedback",
	Short: "Generate a feedback schedule for unit cadets.",
	Long:  ``,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if fbkOutfile == "" {
			fbkOutfile = "FeedbackSchedule.pdf"
		}

		err := feedback.BuildSchedule(fbkOutfile, fbkMbrReport)
		if err != nil {
			logging.Error().Err(err).Msg("Failed to generate feedback schedule")
			os.Exit(1)
		}
	},
}

func init() {
	attendanceCmd.Flags().StringVarP(&fbkOutfile, "out", "o", "", "output file path (defaults to the log date)")
	attendanceCmd.Flags().StringVarP(&fbkMbrReport, "membership-report", "r", "", "file path to eServices Membership report (skips CAPWATCH access)")

	rootCmd.AddCommand(feedbackCmd)
}
