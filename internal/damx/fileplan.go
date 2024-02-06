package damx

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/spf13/cobra"

	"github.com/ut080/bcs-portal/internal/fileplan"
	"github.com/ut080/bcs-portal/internal/logging"
)

var csvOutFileStr string
var pdfOutFileStr string

var fileplanCmd = &cobra.Command{
	Use:   "fileplan [FILE_PLAN_YAML]",
	Short: "Read the YAML file and generate a CSV for creating file plan labels as well as a PDF file plan. ",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.Logger{}
		filePlanCfg, err := files.NewFile(args[0], logger.DefaultLogger())
		if err != nil {
			logging.Error().Err(err).Str("filePlanCfg", args[0]).Msg("Failed to acquire reference for file plan config")
			os.Exit(1)
		}

		var csvOutFile files.File
		if csvOutFileStr == "" {
			csvOutFileStr = filepath.Join(filePlanCfg.Dir(), fmt.Sprintf("%s.csv", filePlanCfg.Base()))
		}
		csvOutFile, err = files.NewFile(csvOutFileStr, logger.DefaultLogger())
		if err != nil {
			logging.Error().Err(err).Str("csvOutFileStr", csvOutFileStr).Msg("Failed to acquire reference for file plan output CSV")
			os.Exit(1)
		}

		var pdfOutFile files.File
		if pdfOutFileStr == "" {
			pdfOutFileStr = filepath.Join(filePlanCfg.Dir(), fmt.Sprintf("%s.pdf", filePlanCfg.Base()))
		}
		pdfOutFile, err = files.NewFile(pdfOutFileStr, logger.DefaultLogger())
		if err != nil {
			logging.Error().Err(err).Str("pdfOutFileStr", pdfOutFileStr).Msg("Failed to acquire reference for file plan output PDF")
			os.Exit(1)
		}

		err = fileplan.BuildFilePlan(filePlanCfg, csvOutFile, pdfOutFile, logger)
		if err != nil {
			logging.Error().Err(err).Msg("Failed to generate file plan")
			os.Exit(1)
		}
	},
}

func init() {
	fileplanCmd.Flags().StringVarP(&csvOutFileStr, "csv", "c", "", "output file path for the CSV file (defaults to the basename of the fileplan config)")
	fileplanCmd.Flags().StringVarP(&csvOutFileStr, "pdf", "p", "", "output file path for the PDF file (defaults to the basename of the fileplan config)")

	RootCmd.AddCommand(fileplanCmd)
}
