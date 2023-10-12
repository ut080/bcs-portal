package damx

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/fileplan"
	"github.com/ut080/bcs-portal/internal/files"
	"github.com/ut080/bcs-portal/internal/logging"
)

var csvOutfile string
var pdfOutfile string

var attendanceCmd = &cobra.Command{
	Use:   "fileplan [FILE_PLAN_YAML]",
	Short: "Read the YAML file and generate a CSV for creating file plan labels as well as a PDF file plan. ",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		planPath, planBase, planExt, err := files.DecomposePath(args[1])

		if csvOutfile == "" {
			csvOutfile = fmt.Sprintf("%s.csv", planBase)
		}

		if pdfOutfile == "" {
			pdfOutfile = fmt.Sprintf("%s.pdf", planBase)
		}

		err := fileplan.BuildFilePlan(fileplanConfig, csvOutfile, pdfOutfile)
		if err != nil {
			logging.Error().Err(err).Msg("Failed to generate file plan")
			os.Exit(1)
		}
	},
}

func init() {
	attendanceCmd.Flags().StringVarP(&csvOutfile, "csv", "c", "", "output file path for the CSV file (defaults to the basename of the fileplan config)")
	attendanceCmd.Flags().StringVarP(&csvOutfile, "pdf", "p", "", "output file path for the PDF file (defaults to the basename of the fileplan config)")

	rootCmd.AddCommand(attendanceCmd)
}
