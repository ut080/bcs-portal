package yaml

import (
	"fmt"

	"github.com/ut080/bcs-portal/pkg/filing"
)

type DispositionRule struct {
	RuleNumber   uint   `yaml:"rule_number"`
	RecordType   string `yaml:"record_type"`
	AutoCutoff   bool   `yaml:"auto_cutoff"`
	Cutoff       string `yaml:"cutoff"`
	DisposeAfter int    `yaml:"dispose_after"`
	Disposition  string `yaml:"disposition"`
}

func (dr DispositionRule) DomainDispositionRule(tableNumber uint) filing.DispositionRule {
	var cutoff filing.Cutoff
	switch dr.Cutoff {
	case "":
		cutoff = filing.NoCutoff
	case "30 Sep":
		cutoff = filing.FiscalYearCutoff
	case "31 Dec":
		cutoff = filing.CalendarYearCutoff
	default:
		panic(fmt.Sprintf("invalid cutoff value parsed from YAML: %s", dr.Cutoff))
	}

	return filing.NewDispositionRule(
		tableNumber,
		dr.RuleNumber,
		dr.RecordType,
		dr.AutoCutoff,
		cutoff,
		dr.DisposeAfter,
		dr.Disposition,
	)
}

type DispositionTable struct {
	TableNumber uint              `yaml:"table_number"`
	Title       string            `yaml:"title"`
	Rules       []DispositionRule `yaml:"rules"`
}

func (dt DispositionTable) DomainDispositionTable() filing.DispositionTable {
	rules := make([]filing.DispositionRule, len(dt.Rules))

	for _, rule := range dt.Rules {
		rules = append(rules, rule.DomainDispositionRule(dt.TableNumber))
	}

	return filing.NewDispositionTable(
		dt.TableNumber,
		dt.Title,
		rules,
	)
}
