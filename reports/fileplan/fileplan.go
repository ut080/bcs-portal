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
	title  string
	table  uint
	rule   uint
}

func NewFilePlanItem(item filing.FilePlanItem) FilePlanItem {
	return FilePlanItem{
		itemID: item.ItemID(),
		title:  item.Title(),
		table:  item.Disposition().TableNumber(),
		rule:   item.Disposition().RuleNumber(),
	}
}

func (fpi FilePlanItem) LaTeX() string {
	if fpi.table == 0 && fpi.rule == 0 {
		return fmt.Sprintf("%s & %s & & \\\\\n", fpi.itemID, fpi.title)
	}

	return fmt.Sprintf("%s & %s & %d & %d \\\\\n", fpi.itemID, fpi.title, fpi.table, fpi.rule)
}

func unpackFilePlanItems(item filing.FilePlanItem) []FilePlanItem {
	var items []FilePlanItem

	items = append(items, NewFilePlanItem(item))

	for _, subitem := range item.Subitems() {
		items = append(items, unpackFilePlanItems(subitem)...)
	}

	return items
}
