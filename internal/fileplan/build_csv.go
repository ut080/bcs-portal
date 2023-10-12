package fileplan

import (
	"fmt"

	"github.com/ut080/bcs-portal/pkg/filing"
)

func filePlanItemAsCsvRow(item filing.FilePlanItem) []string {
	row := make([]string, 4)

	row[0] = item.ItemID()
	row[1] = item.Title()
	row[2] = fmt.Sprintf("%s, %s", item.Table(), item.Rule())
	row[3] = item.Disposition().Instructions()

	return row
}

func filePlanSubItemsToRows(item filing.FilePlanItem) [][]string {
	var rows [][]string
	rows = append(rows, filePlanItemAsCsvRow(item))

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
