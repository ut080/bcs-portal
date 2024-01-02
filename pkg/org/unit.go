package org

import (
	"github.com/google/uuid"
)

type Unit struct {
	id            uuid.UUID
	charterNumber string
	unitType      UnitType
	name          string
	commander     DutyAssignment
}

func NewUnit(
	id uuid.UUID,
	charterNumber string,
	unitType UnitType,
	name string,
	commander DutyAssignment,
) Unit {
	return Unit{
		id:            id,
		charterNumber: charterNumber,
		unitType:      unitType,
		name:          name,
		commander:     commander,
	}
}

func (u Unit) ID() uuid.UUID {
	return u.id
}

func (u Unit) CharterNumber() string {
	return u.charterNumber
}

func (u Unit) UnitType() UnitType {
	return u.unitType
}

func (u Unit) Name() string {
	return u.name
}

func (u Unit) Commander() DutyAssignment {
	return u.commander
}
