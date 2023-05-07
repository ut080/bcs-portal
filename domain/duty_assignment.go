package domain

import (
	"sort"
)

type DutyAssignment struct {
	OfficeSymbol string
	DutyTitle    string
	Assistant    bool
	Acting       bool
	MinGrade     *Grade
	MaxGrade     *Grade
	Assignee     *Member
}

func SortDutyAssignmentsByAssigneeName(da []DutyAssignment) {
	sort.Slice(da, func(i, j int) bool {
		if da[i].Assignee == nil {
			return false
		}

		if da[j].Assignee == nil {
			return true
		}

		if da[i].Assignee.LastName == da[j].Assignee.LastName {
			return da[i].Assignee.FirstName < da[j].Assignee.FirstName
		}

		return da[i].Assignee.LastName < da[j].Assignee.LastName
	})
}
