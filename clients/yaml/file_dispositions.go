package yaml

import (
	"github.com/ut080/bcs-portal/pkg/filing"
)

type DispositionRule struct {
	RuleNumber  uint   `yaml:"rule_number"`
	RecordType  string `yaml:"record_type"`
	Cutoff      string `yaml:"cutoff"`
	Disposition string `yaml:"disposition"`
}

func (dr DispositionRule) DomainDispositionRule(tableNumber uint) filing.DispositionRule {
	return filing.NewDispositionRule(
		tableNumber,
		dr.RuleNumber,
		dr.RecordType,
		dr.Cutoff,
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
