package fileplan

import (
	"fmt"
	"strings"
	"time"

	"github.com/ut080/bcs-portal/pkg/filing"
)

const (
	filePlanPreamble = `\documentclass[12pt]{article}

\usepackage{fancyhdr}
\usepackage[
    top=1.0in,
    bottom=0.75in,
    left=0.25in,
    right=0.25in,
    headheight=15pt
]{geometry}
\usepackage{lastpage}
\usepackage{xltabular}

\fancyhf{}
\renewcommand{\headrulewidth}{0pt}
\fancyhead[L]{Name: %{PREPARER}}
\fancyhead[R]{Date Prepared: %{DATE_PREPARED}}
\fancyfoot[L]{%{TITLE}}
\fancyfoot[R]{Page\ \thepage\ of\ \pageref{LastPage}}
\pagestyle{fancy}

\begin{document}
    \noindent
    \begin{xltabular}{\textwidth}{lXrr}
        \underline{Item} & \underline{Title} & \underline{Table} & \underline{Rule} \\
        \endhead

`
	filePlanEndTemplate = `    \end{xltabular}
\end{document}`
)

var empty FilePlanItem = FilePlanItem{}

type FilePlan struct {
	planTitle string
	preparer  string
	prepared  time.Time
	items     []FilePlanItem
}

func NewFilePlan(plan filing.FilePlan) FilePlan {
	var items []FilePlanItem

	for _, item := range plan.Items() {
		items = append(items, unpackFilePlanItems(item)...)
	}

	items = chompLines(items)

	return FilePlan{
		planTitle: plan.PlanTitle(),
		preparer:  plan.Preparer(),
		prepared:  plan.Prepared(),
		items:     items,
	}
}

func (fp FilePlan) LaTeX() string {
	latex := filePlanPreamble

	latex = strings.Replace(latex, "%{PREPARER}", fp.preparer, -1)
	latex = strings.Replace(latex, "%{DATE_PREPARED}", fp.prepared.Format("2 Jan 2006"), -1)
	latex = strings.Replace(latex, "%{TITLE}", fp.planTitle, -1)

	for _, item := range fp.items {
		latex += item.LaTeX()
	}

	latex += filePlanEndTemplate

	return latex
}

type FilePlanItem struct {
	itemID string
	level  int
	title  string
	table  uint
	rule   uint
}

func NewFilePlanItem(item filing.FilePlanItem) FilePlanItem {
	return FilePlanItem{
		itemID: item.ItemID(),
		level:  item.Level(),
		title:  item.Title(),
		table:  item.Disposition().TableNumber(),
		rule:   item.Disposition().RuleNumber(),
	}
}

func (fpi FilePlanItem) LaTeX() string {
	itemID := fpi.itemID
	title := fpi.title
	var table string
	var rule string

	if fpi.table != 0 && fpi.rule != 0 {
		table = fmt.Sprintf("%d", fpi.table)
		rule = fmt.Sprintf("%d", fpi.rule)
	}

	if fpi.level == -1 {
		itemID = fmt.Sprintf("\\textbf{\\underline{%s}}", itemID)
		title = fmt.Sprintf("\\textbf{\\underline{%s}}", title)
	} else if fpi.level > 1 {
		itemID = fmt.Sprintf("\\textbf{%s}", itemID)
		title = fmt.Sprintf("\\textbf{%s}", title)
	} else if fpi.level == 1 {
		itemID = fmt.Sprintf("\\textit{%s}", itemID)
		title = fmt.Sprintf("\\textit{%s}", title)
	}

	return fmt.Sprintf("%s & %s & %s & %s \\\\\n", itemID, title, table, rule)
}

func unpackFilePlanItems(item filing.FilePlanItem) []FilePlanItem {
	var items []FilePlanItem

	switch item.Level() {
	case 0:
		// File level 0 has no subitems, so we just need to return this item
		return append(items, NewFilePlanItem(item))
	case 1:
		// File Level 1 is grouped with its subitems with a blank line at the end
		items = append(items, NewFilePlanItem(item))
		for _, subitem := range item.Subitems() {
			items = append(items, unpackFilePlanItems(subitem)...)
		}
		items = append(items, FilePlanItem{})
	default:
		// All other file levels have a blank line before, between the item and its subitems, and a blank line at the end of the group
		items = append(items, FilePlanItem{}, NewFilePlanItem(item), FilePlanItem{})
		for _, subitem := range item.Subitems() {
			items = append(items, unpackFilePlanItems(subitem)...)
		}
		items = append(items, FilePlanItem{})
	}

	return items
}

// chompLines removes excess empty FilePlanItems.
func chompLines(items []FilePlanItem) []FilePlanItem {
	var chompedItems []FilePlanItem

	var previousEmpty bool
	for _, item := range items {
		if item == empty {
			if previousEmpty {
				continue
			}
			previousEmpty = true
		} else {
			previousEmpty = false
		}

		chompedItems = append(chompedItems, item)
	}

	return chompedItems
}
