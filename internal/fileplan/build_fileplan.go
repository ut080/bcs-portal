package fileplan

import (
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/clients/yaml"
	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/filing"
)

func loadFileDispositionRules(logger logging.Logger) (map[uint]filing.DispositionTable, error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to access config directory")
	}

	dispRulesCfg := make(map[string]yaml.DispositionTable)
	dispRulesCfgPath := filepath.Join(cfgDir, "cfg", "disposition_instructions.yaml")
	err = yaml.LoadFromFile(dispRulesCfgPath, &dispRulesCfg, logger)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dispRules := make(map[uint]filing.DispositionTable)
	for _, table := range dispRulesCfg {
		dispRules[table.TableNumber] = table.DomainDispositionTable()
	}

	return dispRules, nil
}

func readFilePlan(input string, dispositionRules map[uint]filing.DispositionTable, logger logging.Logger) (filing.FilePlan, error) {
	var filePlanCfg yaml.FilePlan
	err := yaml.LoadFromFile(input, &filePlanCfg, logger)
	if err != nil {
		return filing.FilePlan{}, errors.WithStack(err)
	}

	return filePlanCfg.DomainFilePlan(dispositionRules, logger), nil
}

func generateCSV(filePlan filing.FilePlan, outCSV string) error {
	csv, err := newCSV(outCSV)
	if err != nil {
		return errors.WithStack(err)
	}

	plan := convertFilePlanToRows(filePlan)

	err = csv.WriteAll(plan)
	if err != nil {
		return errors.WithStack(err)
	}

	csv.Flush()

	return nil
}

func generatePDF(fileplan filing.FilePlan, outPDF string) error {
}

func BuildFilePlan(input, outCSV, outPDF string) error {
	logger := logging.Logger{}

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

	return nil
}
