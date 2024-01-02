package org

import (
	"github.com/google/uuid"
)

type StaffGroup struct {
	id        uuid.UUID
	name      string
	subgroups []StaffGroup
	leader    DutyAssignment
}

func NewStaffGroup(
	id uuid.UUID,
	name string,
	subgroups []StaffGroup,
	leader DutyAssignment,
) StaffGroup {
	return StaffGroup{
		id:        id,
		name:      name,
		subgroups: subgroups,
		leader:    leader,
	}
}

func (sg StaffGroup) ID() uuid.UUID {
	return sg.id
}

func (sg StaffGroup) Name() string {
	return sg.name
}

func (sg StaffGroup) Subgroups() []StaffGroup {
	return sg.subgroups
}

func (sg StaffGroup) Leader() DutyAssignment {
	return sg.leader
}
