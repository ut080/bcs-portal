package org

import (
	"github.com/google/uuid"
)

type Flight struct {
	id              uuid.UUID
	name            string
	abbreviation    string
	flightCommander DutyAssignment
	flightSergeant  DutyAssignment
	elements        []Element
}

func NewFlight(
	id uuid.UUID,
	name string,
	abbreviation string,
	flightCommander DutyAssignment,
	flightSergeant DutyAssignment,
	elements []Element,
) Flight {
	return Flight{
		id:              id,
		name:            name,
		abbreviation:    abbreviation,
		flightCommander: flightCommander,
		flightSergeant:  flightSergeant,
		elements:        elements,
	}
}

func (f Flight) ID() uuid.UUID {
	return f.id
}

func (f Flight) Name() string {
	return f.name
}

func (f Flight) Abbreviation() string {
	return f.abbreviation
}

func (f Flight) FlightCommander() DutyAssignment {
	return f.flightCommander
}

func (f Flight) FlightSergeant() DutyAssignment {
	return f.flightSergeant
}

func (f Flight) Elements() []Element {
	return f.elements
}
