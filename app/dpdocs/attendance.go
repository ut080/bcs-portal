package dpdocs

import (
	"github.com/spf13/cobra"
)

var capwatchPassword string
var outfile string

var attendanceCmd = &cobra.Command{
	Use:   "attendance [TO DEF FILE]",
	Short: "Generate a barcode attendance log from CAPWATCH and unit TO data.",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	attendanceCmd.Flags().StringVarP(&capwatchPassword, "password", "p", "", "eServices password")
	attendanceCmd.Flags().StringVarP(&outfile, "out", "o", "attendance.pdf", "output file path")

	rootCmd.AddCommand(attendanceCmd)
}
