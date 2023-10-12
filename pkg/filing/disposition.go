package filing

import (
	"fmt"
)

type InvalidRuleNumber struct {
	TableNumber   uint
	TableTitle    string
	BadRuleNumber uint
}

func (irn InvalidRuleNumber) Error() string {
	return fmt.Sprintf("invalid rule ruleNumber on table %s (%s): %d", irn.TableNumber, irn.TableTitle, irn.BadRuleNumber)
}

type DispositionTable struct {
	number uint
	title  string
	rules  map[uint]DispositionRule
}

func NewDispositionTable(number uint, title string, rules []DispositionRule) DispositionTable {
	r := make(map[uint]DispositionRule)

	for _, v := range rules {
		r[v.ruleNumber] = v
	}

	return DispositionTable{
		number: number,
		title:  title,
		rules:  r,
	}
}

func (dt DispositionTable) Number() uint {
	return dt.number
}

func (dt DispositionTable) Title() string {
	return dt.title
}

func (dt DispositionTable) Rule(number uint) (DispositionRule, error) {
	rule, ok := dt.rules[number]
	if !ok {
		return DispositionRule{}, InvalidRuleNumber{
			TableNumber:   dt.number,
			TableTitle:    dt.title,
			BadRuleNumber: number,
		}
	}

	return rule, nil
}

type DispositionRule struct {
	tableNumber  uint
	ruleNumber   uint
	recordType   string
	cutoff       string
	instructions string
}

func NewDispositionRule(tableNumber, ruleNumber uint, recordType, cutoff, instructions string) DispositionRule {
	return DispositionRule{
		tableNumber:  tableNumber,
		ruleNumber:   ruleNumber,
		recordType:   recordType,
		cutoff:       cutoff,
		instructions: instructions,
	}
}

func (dr DispositionRule) TableNumber() uint {
	return dr.tableNumber
}

func (dr DispositionRule) RuleNumber() uint {
	return dr.ruleNumber
}

func (dr DispositionRule) RecordType() string {
	return dr.recordType
}

func (dr DispositionRule) Cutoff() string {
	return dr.cutoff
}

func (dr DispositionRule) Instructions() string {
	return dr.instructions
}
