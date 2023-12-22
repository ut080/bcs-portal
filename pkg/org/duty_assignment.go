package org

import (
	"github.com/google/uuid"
)

type DutyAssignment struct {
	id           uuid.UUID
	Title        string
	OfficeSymbol string
	Assistant    bool
	DutyTitle    DutyTitle
	Assignee     *Member
}

func NewDutyAssignment(
	id uuid.UUID,
	title string,
	officeSymbol string,
	assistant bool,
	dutyTitle DutyTitle,
	assignee *Member,
) DutyAssignment {
	return DutyAssignment{
		id:           id,
		Title:        title,
		OfficeSymbol: officeSymbol,
		Assistant:    assistant,
		DutyTitle:    dutyTitle,
		Assignee:     assignee,
	}
}

func (da DutyAssignment) ID() uuid.UUID {
	return da.id
}
