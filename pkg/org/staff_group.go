package org

import (
	"github.com/google/uuid"
)

type StaffGroup struct {
	id        uuid.UUID
	name      string
	subGroups []StaffSubgroup
}

func NewStaffGroup(id uuid.UUID, name string, subgroups []StaffSubgroup) StaffGroup {
	return StaffGroup{
		id:        id,
		name:      name,
		subGroups: subgroups,
	}
}

func (sg StaffGroup) ID() uuid.UUID {
	return sg.id
}

func (sg StaffGroup) Name() string {
	return sg.name
}

func (sg StaffGroup) SubGroups() []StaffSubgroup {
	return sg.subGroups
}
