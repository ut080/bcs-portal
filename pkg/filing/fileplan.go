package filing

import (
	"fmt"
	"time"
)

type FilePlan struct {
	planTitle string
	preparer  string
	prepared  time.Time
	items     []FilePlanItem
}

func NewFilePlan(planTitle, preparer string, prepared time.Time, items []FilePlanItem) FilePlan {
	return FilePlan{
		planTitle: planTitle,
		preparer:  preparer,
		prepared:  prepared,
		items:     items,
	}
}

func (fp FilePlan) PlanTitle() string {
	return fp.planTitle
}

func (fp FilePlan) Preparer() string {
	return fp.preparer
}

func (fp FilePlan) Prepared() time.Time {
	return fp.prepared
}

func (fp FilePlan) Items() []FilePlanItem {
	return fp.items
}

type FilePlanItem struct {
	itemID   string
	title    string
	rule     DispositionRule
	subitems []FilePlanItem
}

func NewFilePlanItem(itemID string, title string, rule DispositionRule, subitems []FilePlanItem) FilePlanItem {
	return FilePlanItem{
		itemID:   itemID,
		title:    title,
		rule:     rule,
		subitems: subitems,
	}
}

func (fpi FilePlanItem) ItemID() string {
	return fpi.itemID
}

func (fpi FilePlanItem) Title() string {
	return fpi.title
}

func (fpi FilePlanItem) Table() uint {
	return fpi.rule.tableNumber
}

func (fpi FilePlanItem) Rule() uint {
	return fpi.rule.ruleNumber
}

func (fpi FilePlanItem) Disposition() DispositionRule {
	return fpi.rule
}

func (fpi FilePlanItem) Subitems() []FilePlanItem {
	return fpi.subitems
}

func (fpi FilePlanItem) String() string {
	return fmt.Sprintf("%s  %s", fpi.itemID, fpi.title)
}
