package filing

import (
	"fmt"
	"time"

	"github.com/ut080/bcs-portal/pkg/org"
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
	itemID     string
	title      string
	short      string
	rule       DispositionRule
	folderType FolderType
	electronic bool
	personnel  org.MemberType
	level      int
	subitems   []FilePlanItem
}

func NewFilePlanItem(itemID string, title, short string, rule DispositionRule, folderType FolderType, electronic bool, personnel org.MemberType, subitems []FilePlanItem, root bool) FilePlanItem {
	var level int
	if root {
		level = -1
	} else if subitems != nil {
		var largest int
		for _, v := range subitems {
			if v.Level() > largest {
				largest = v.Level()
			}
		}

		level = largest + 1
	}

	return FilePlanItem{
		itemID:     itemID,
		title:      title,
		short:      short,
		rule:       rule,
		folderType: folderType,
		electronic: electronic,
		personnel:  personnel,
		level:      level,
		subitems:   subitems,
	}
}

func (fpi FilePlanItem) ItemID() string {
	return fpi.itemID
}

func (fpi FilePlanItem) Title() string {
	return fpi.title
}

func (fpi FilePlanItem) ShortTitle() string {
	return fpi.short
}

func (fpi FilePlanItem) Table() uint {
	return fpi.rule.tableNumber
}

func (fpi FilePlanItem) Rule() uint {
	return fpi.rule.ruleNumber
}

func (fpi FilePlanItem) FolderType() FolderType {
	return fpi.folderType
}

func (fpi FilePlanItem) Electronic() bool {
	return fpi.electronic
}

func (fpi FilePlanItem) Disposition() DispositionRule {
	return fpi.rule
}

func (fpi FilePlanItem) Level() int {
	return fpi.level
}

func (fpi FilePlanItem) Subitems() []FilePlanItem {
	return fpi.subitems
}

func (fpi FilePlanItem) String() string {
	return fmt.Sprintf("%s  %s", fpi.itemID, fpi.title)
}
