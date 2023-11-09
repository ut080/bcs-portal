package filing

import (
	"fmt"
	"time"
)

type InvalidRuleNumber struct {
	TableNumber   uint
	TableTitle    string
	BadRuleNumber uint
}

func (irn InvalidRuleNumber) Error() string {
	return fmt.Sprintf("invalid rule ruleNumber on table %s (%s): %d", irn.TableNumber, irn.TableTitle, irn.BadRuleNumber)
}

type Cutoff int

const (
	NoCutoff Cutoff = iota
	FiscalYearCutoff
	CalendarYearCutoff
)

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
	// tableNumber represents the table at the end of CAPR 10-2 that defines this rule.
	tableNumber uint

	// ruleNumber represents the rule number in the given CAPR 10-2 table.
	ruleNumber uint

	// recordType is the description of the records to which this rule is supposed to be applied, as given in CAPR 10-2.
	recordType string

	// autoCutoff is the flag that indicates whether the damx tool should automatically rotate/destroy files to which this rule applies.
	autoCutoff bool

	// cutoff is an enum value. Valid values are NoCutoff, FiscalYearCutoff (representing CAPR 10-2 cutoff dates of 30 Sep)
	// or CalendarYearCutoff (representing CAPR 10-2 cutoff dates of 31 Dec).
	cutoff Cutoff

	// disposeAfter is the number of years after the cutoff when this record should be destroyed. The value of 0 is used
	// wherever CAPR 10-2 gives a disposition instruction to the effect of "dispose when no longer needed," thus the
	// file will need to be manually deleted. -1 is used wherever CAPR 10-2 dictates that a record should be "retained permanently."
	disposeAfter int

	// instructions represent the disposition instructions given by CAPR 10-2.
	instructions string
}

func NewDispositionRule(tableNumber, ruleNumber uint, recordType string, autoCutoff bool, cutoff Cutoff, disposeAfter int, instructions string) DispositionRule {
	// The autoCutoff flag is not valid if a cutoff type is set, or if the disposeAfter value is less than 1.
	if !autoCutoff && (cutoff == NoCutoff || disposeAfter < 1) {

	}

	return DispositionRule{
		tableNumber:  tableNumber,
		ruleNumber:   ruleNumber,
		recordType:   recordType,
		autoCutoff:   autoCutoff,
		cutoff:       cutoff,
		disposeAfter: disposeAfter,
		instructions: instructions,
	}
}

func (dr DispositionRule) Empty() bool {
	return dr.tableNumber == 0 || dr.ruleNumber == 0
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

func (dr DispositionRule) AutoCutoff() bool {
	return dr.autoCutoff
}

func (dr DispositionRule) Cutoff() Cutoff {
	return dr.cutoff
}

// CutoffDate determines the date this rule defines as the file's cutoff based on the given calendar year.
// Note that if this method is called after midnight on 1 October (local time), the year is incremented before being
// returned. This represents the cutoff being at the end of the current fiscal year.
func (dr DispositionRule) CutoffDate(year int) time.Time {
	switch dr.cutoff {
	case NoCutoff:
		return time.Time{}
	case FiscalYearCutoff:
		if time.Now().After(time.Date(year, 10, 1, 0, 0, 0, 0, time.Local)) {
			return time.Date(year+1, 9, 30, 0, 0, 0, 0, time.Local)
		}
		return time.Date(year, 9, 30, 0, 0, 0, 0, time.Local)
	case CalendarYearCutoff:
		return time.Date(year, 12, 31, 0, 0, 0, 0, time.Local)
	default:
		panic(fmt.Sprintf("invalid Cutoff value: %d", dr.cutoff))
	}
}

func (dr DispositionRule) DisposeAfter() int {
	return dr.disposeAfter
}

func (dr DispositionRule) Instructions() string {
	return dr.instructions
}
