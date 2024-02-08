package damx

import (
	"fmt"
	"os"
	"time"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/attendance"
	"github.com/ut080/bcs-portal/internal/logging"
)

var attOutfileStr string

var attendanceCmd = &cobra.Command{
	Use:   "attendance [LogDate]",
	Short: "Generate a barcode attendance log.",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run:   runAttendance,
}

func runAttendance(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

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

	err = attendance.BuildBarcodeLog(attOutfile, logDate, logger)
	if err != nil {
		logging.Error().Err(err).Msg("Failed to generate barcode attendance log")
		os.Exit(1)
	}
}

func init() {
	attendanceCmd.Flags().StringVarP(&attOutfileStr, "out", "o", "", "output file path (defaults to the log date)")

	RootCmd.AddCommand(attendanceCmd)
}
