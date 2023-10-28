package yaml

import (
	"fmt"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/filing"
)

type FilePlan struct {
	PlanTitle string         `yaml:"plan_title"`
	Preparer  string         `yaml:"preparer"`
	Prepared  Date           `yaml:"prepared"`
	Items     []FilePlanItem `yaml:"items"`
}

func (fp FilePlan) DomainFilePlan(dispositionRules map[uint]filing.DispositionTable, logger logging.Logger) filing.FilePlan {
	var items []filing.FilePlanItem
	for i, item := range fp.Items {
		items = append(items, item.DomainFilePlanItem(fmt.Sprintf("%d.", i+1), dispositionRules, logger))
	}

	return filing.NewFilePlan(fp.PlanTitle, fp.Preparer, fp.Prepared.Time, items)
}

type FilePlanItem struct {
	Title         string         `yaml:"title"`
	Table         uint           `yaml:"table"`
	Rule          uint           `yaml:"rule"`
	DontMakeLabel bool           `yaml:"dont_make_label"`
	Subitems      []FilePlanItem `yaml:"subitems"`
}

func (fpi FilePlanItem) DomainFilePlanItem(itemID string, dispositionRules map[uint]filing.DispositionTable, logger logging.Logger) filing.FilePlanItem {
	var subitems []filing.FilePlanItem

	if fpi.Subitems != nil {
		for i, subitem := range fpi.Subitems {
			subitemID := fmt.Sprintf("%s%d.", itemID, i+1)
			subitems = append(subitems, subitem.DomainFilePlanItem(subitemID, dispositionRules, logger))
		}
	}

	var table filing.DispositionTable
	if fpi.Table != 0 {
		var ok bool
		table, ok = dispositionRules[fpi.Table]
		if !ok {
			logging.Warn().Str("item_id", itemID).Uint("table", fpi.Table).Msg(
				"invalid disposition table, item will be parsed with no disposition rules",
			)

			table = filing.DispositionTable{}
		}
	}

	var rule filing.DispositionRule
	if table.Number() != 0 {
		var err error
		rule, err = table.Rule(fpi.Rule)
		if err != nil {
			logging.Warn().Str("item_id", itemID).Uint("table", fpi.Table).Uint("rule", fpi.Rule).Msg(
				"invalid disposition rule, item will be parsed with no disposition rules",
			)

		}

	}

	return filing.NewFilePlanItem(itemID, fpi.Title, rule, fpi.DontMakeLabel, subitems)
}
