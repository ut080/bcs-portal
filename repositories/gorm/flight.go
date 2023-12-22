package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type Flight struct {
	ID                uuid.UUID
	Name              string
	Abbreviation      string
	FlightCommanderID uuid.UUID
	FlightSergeantID  uuid.UUID
	FlightCommander   DutyAssignment
	FlightSergeant    DutyAssignment
	Elements          []Element
}

func (f *Flight) FromDomainObject(object pkg.DomainObject) error {
	v, ok := object.(org.Flight)
	if !ok {
		return errors.New("not a valid domain Flight object")
	}

	flightCommander := &DutyAssignment{}
	err := flightCommander.FromDomainObject(v.FlightCommander)
	if err != nil {
		return errors.WithStack(err)
	}

	flightSergeant := &DutyAssignment{}
	err = flightSergeant.FromDomainObject(v.FlightSergeant)
	if err != nil {
		return errors.WithStack(err)
	}

	var elements []Element
	for _, e := range v.Elements {
		element := Element{FlightID: f.ID}
		err = element.FromDomainObject(e)
		if err != nil {
			return errors.WithStack(err)
		}

		elements = append(elements, element)
	}

	f.ID = v.ID()
	f.Name = v.Name
	f.Abbreviation = v.Abbreviation
	f.FlightCommanderID = v.FlightCommander.ID()
	f.FlightSergeantID = v.FlightSergeant.ID()
	f.Elements = elements

	return nil
}

func (f *Flight) ToDomainObject() pkg.DomainObject {
	flightCommander := f.FlightCommander.ToDomainObject().(org.DutyAssignment)
	flightSergeant := f.FlightSergeant.ToDomainObject().(org.DutyAssignment)

	var elements []org.Element
	for _, v := range f.Elements {
		elem := v.ToDomainObject().(org.Element)
		elements = append(elements, elem)
	}

	return org.NewFlight(f.ID, f.Name, f.Abbreviation, flightCommander, flightSergeant, elements)
}
