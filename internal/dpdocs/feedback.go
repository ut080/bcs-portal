package dpdocs

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/feedback"
	"github.com/ut080/bcs-portal/internal/logging"
)

var feedbackCmd = &cobra.Command{
	Use:   "feedback",
	Short: "Generate a feedback schedule for unit cadets.",
	Long:  ``,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := feedback.BuildSchedule(outfile, mbrReport)
		if err != nil {
			logging.Error().Err(err).Msg("Failed to generate feedback schedule")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(feedbackCmd)
}
