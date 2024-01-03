package org

import (
	"github.com/google/uuid"
)

type DutyAssignment struct {
	id           uuid.UUID
	title        string
	officeSymbol string
	assistant    bool
	dutyTitle    DutyTitle
	assignee     *Member
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
		title:        title,
		officeSymbol: officeSymbol,
		assistant:    assistant,
		dutyTitle:    dutyTitle,
		assignee:     assignee,
	}
}

func (da DutyAssignment) ID() uuid.UUID {
	return da.id
}

func (da DutyAssignment) Title() string {
	if da.title != "" {
		return da.title
	}

	return da.dutyTitle.Title()
}

func (da DutyAssignment) OfficeSymbol() string {
	return da.officeSymbol
}

func (da DutyAssignment) Assistant() bool {
	return da.assistant
}

func (da DutyAssignment) DutyTitle() DutyTitle {
	return da.dutyTitle
}

func (da DutyAssignment) Assignee() *Member {
	return da.assignee
}
