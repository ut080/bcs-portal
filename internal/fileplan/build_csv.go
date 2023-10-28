package fileplan

import (
	"fmt"
	"unicode"

	"github.com/ut080/bcs-portal/pkg/filing"
)

func filePlanItemAsCsvRow(item filing.FilePlanItem) []string {
	row := make([]string, 4)

	row[0] = item.ItemID()
	row[1] = item.Title()

	// Process LaTeX commands for en- and em-dashes in the title row

	if !item.Disposition().Empty() {
		row[2] = fmt.Sprintf("T%d, R%d", item.Table(), item.Rule())

		if item.Disposition().Cutoff() == "" {
			instructions := []rune(item.Disposition().Instructions())
			instructions[0] = unicode.ToUpper(instructions[0])
			row[3] = string(instructions)
		} else {
			row[3] = fmt.Sprintf("Cut Off: %s/%s", item.Disposition().Cutoff(), item.Disposition().Instructions())
		}
	}

	return row
}

func filePlanSubItemsToRows(item filing.FilePlanItem) [][]string {
	var rows [][]string

	if !item.DontMakeLabel() {
		rows = append(rows, filePlanItemAsCsvRow(item))
	}

	for _, subitem := range item.Subitems() {
		rows = append(rows, filePlanSubItemsToRows(subitem)...)
	}

	return rows
}

func convertFilePlanToRows(filePlan filing.FilePlan) [][]string {
	var rows [][]string

	for _, item := range filePlan.Items() {
		rows = append(rows, filePlanSubItemsToRows(item)...)
	}

	return rows
}
