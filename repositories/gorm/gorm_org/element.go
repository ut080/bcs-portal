package gorm_org

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type Element struct {
	ID                  uuid.UUID
	Name                string
	FlightID            uuid.UUID
	ElementLeaderID     uuid.UUID
	AsstElementLeaderID uuid.UUID
	ElementLeader       DutyAssignment
	AsstElementLeader   DutyAssignment
	Members             []ElementMember
}

func (e *Element) FromDomainObject(object pkg.DomainObject) error {
	if e.FlightID == uuid.Nil {
		return errors.New("attempt to call Element.FromDomainObject without first setting FlightID")
	}

	elem, ok := object.(org.Element)
	if !ok {
		return errors.New("attempt to pass non-org.Element object to Element.FromDomainObject")
	}

	el := DutyAssignment{}
	err := el.FromDomainObject(elem.ElementLeader())
	if err != nil {
		return errors.WithStack(err)
	}

	ael := DutyAssignment{}
	err = ael.FromDomainObject(elem.ElementLeader())
	if err != nil {
		return errors.WithStack(err)
	}

	var members []ElementMember
	for _, v := range elem.Members() {
		mbr := ElementMember{ElementID: elem.ID()}
		err = mbr.FromDomainObject(v)
		if err != nil {
			return errors.WithStack(err)
		}

		members = append(members, mbr)
	}

	e.ID = elem.ID()
	e.Name = elem.Name()
	e.ElementLeaderID = elem.ElementLeader().ID()
	e.AsstElementLeaderID = elem.AsstElementLeader().ID()
	e.ElementLeader = el
	e.AsstElementLeader = el
	e.Members = members

	return nil
}

func (e *Element) ToDomainObject() pkg.DomainObject {
	obj := e.ElementLeader.ToDomainObject()
	el := obj.(org.DutyAssignment)

	obj = e.AsstElementLeader.ToDomainObject()
	ael := obj.(org.DutyAssignment)

	var members []org.Member
	for _, v := range e.Members {
		obj = v.ToDomainObject()
		mbr := obj.(org.Member)
		members = append(members, mbr)
	}

	return org.NewElement(
		e.ID,
		e.Name,
		el,
		ael,
		members,
	)
}

func (e *Element) BeforeCreate(tx *gorm.DB) error {
	if e.ID != uuid.Nil {
		return nil
	}

	var err error
	e.ID, err = uuid.NewRandom()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
