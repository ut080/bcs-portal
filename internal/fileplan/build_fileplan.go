package fileplan

import (
	"encoding/csv"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/clients/yaml"
	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/filing"
	"github.com/ut080/bcs-portal/reports"
	"github.com/ut080/bcs-portal/reports/fileplan"
)

func loadFileDispositionRules(logger logging.Logger) (map[uint]filing.DispositionTable, error) {
	cfgDir, err := config.CfgDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to access config directory")
	}

	dispRulesCfg := make(map[string]yaml.DispositionTable)
	dispRulesCfgFile, err := files.NewFile(filepath.Join(cfgDir, "cfg", "defs", "disposition_instructions.yaml"), logger.DefaultLogger())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// TODO: Add schema validation
	err = yaml.LoadFromFile(dispRulesCfgFile, &dispRulesCfg, nil, logger)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dispRules := make(map[uint]filing.DispositionTable)
	for _, table := range dispRulesCfg {
		dispRules[table.TableNumber] = table.DomainDispositionTable()
	}

	return dispRules, nil
}

func readFilePlan(planFile files.File, dispositionRules map[uint]filing.DispositionTable, logger logging.Logger) (filing.FilePlan, error) {
	var filePlanCfg yaml.FilePlan

	// TODO: Add schema validation
	err := yaml.LoadFromFile(planFile, &filePlanCfg, nil, logger)
	if err != nil {
		return filing.FilePlan{}, errors.WithStack(err)
	}

	return filePlanCfg.DomainFilePlan(dispositionRules, logger), nil
}

func generateCSV(filePlan filing.FilePlan, outCSV files.File) error {
	f, err := outCSV.Create()
	if err != nil {
		return errors.WithStack(err)
	}
	csvFile := csv.NewWriter(f)

	plan := convertFilePlanToRows(filePlan)

	err = csvFile.WriteAll(plan)
	if err != nil {
		return errors.WithStack(err)
	}

	csvFile.Flush()

	return nil
}

func generatePDF(filePlan filing.FilePlan, outPDF files.File, logger logging.Logger) error {
	plan := fileplan.NewFilePlan(filePlan)

	compiler, err := reports.ConfigureLaTeXCompiler(logger)

	err = compiler.GenerateLaTeX(plan, outPDF, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.CompileLaTeX(outPDF)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func BuildFilePlan(input, outCSV, outPDF files.File, logger logging.Logger) error {
	dispositionRules, err := loadFileDispositionRules(logger)
	if err != nil {
		return errors.WithStack(err)
	}

	filePlan, err := readFilePlan(input, dispositionRules, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = generateCSV(filePlan, outCSV)
	if err != nil {
		return errors.WithStack(err)
	}

	err = generatePDF(filePlan, outPDF, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
