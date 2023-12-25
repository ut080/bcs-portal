package org

import (
	"github.com/google/uuid"
)

type Element struct {
	id                uuid.UUID
	name              string
	elementLeader     DutyAssignment
	asstElementLeader DutyAssignment
	members           []Member
}

func NewElement(
	id uuid.UUID,
	name string,
	elementLeader DutyAssignment,
	asstElementLeader DutyAssignment,
	members []Member,
) Element {
	return Element{
		id:                id,
		name:              name,
		elementLeader:     elementLeader,
		asstElementLeader: asstElementLeader,
		members:           members,
	}
}

func (e Element) ID() uuid.UUID {
	return e.id
}

func (e Element) Name() string {
	return e.name
}

func (e Element) ElementLeader() DutyAssignment {
	return e.elementLeader
}

func (e Element) AsstElementLeader() DutyAssignment {
	return e.asstElementLeader
}

func (e Element) Members() []Member {
	return e.members
}
