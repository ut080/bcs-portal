package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

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
	flt, ok := object.(org.Flight)
	if !ok {
		return errors.New("attempt to pass non-org.Flight object to Flight.FromDomainObject")
	}

	fltCC := DutyAssignment{}
	err := fltCC.FromDomainObject(flt.FlightCommander())
	if err != nil {
		return errors.WithStack(err)
	}

	fltCCF := DutyAssignment{}
	err = fltCCF.FromDomainObject(flt.FlightSergeant())
	if err != nil {
		return errors.WithStack(err)
	}

	var elements []Element
	for _, v := range flt.Elements() {
		elem := Element{FlightID: f.ID}
		err = elem.FromDomainObject(v)
		if err != nil {
			return errors.WithStack(err)
		}

		elements = append(elements, elem)
	}

	f.ID = flt.ID()
	f.Name = flt.Name()
	f.Abbreviation = flt.Abbreviation()
	f.FlightCommanderID = flt.FlightCommander().ID()
	f.FlightSergeantID = flt.FlightSergeant().ID()
	f.FlightCommander = fltCC
	f.FlightSergeant = fltCCF
	f.Elements = elements

	return nil
}

func (f *Flight) ToDomainObject() pkg.DomainObject {
	obj := f.FlightCommander.ToDomainObject()
	fltCC := obj.(org.DutyAssignment)

	obj = f.FlightSergeant.ToDomainObject()
	fltCCF := obj.(org.DutyAssignment)

	var elements []org.Element
	for _, v := range f.Elements {
		e := v.ToDomainObject()
		elem := e.(org.Element)
		elements = append(elements, elem)
	}

	return org.NewFlight(
		f.ID,
		f.Name,
		f.Abbreviation,
		fltCC,
		fltCCF,
		elements,
	)
}

func (f *Flight) BeforeCreate(tx *gorm.DB) error {
	if f.ID != uuid.Nil {
		return nil
	}

	var err error
	f.ID, err = uuid.NewRandom()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
