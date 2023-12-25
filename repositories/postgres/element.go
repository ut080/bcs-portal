package postgres

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

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
		return errors.New("attempt to convert Element domain object without setting FlightID")
	}

	v, ok := object.(org.Element)
	if !ok {
		return errors.New("not a valid domain Element object")
	}

	elementLeader := DutyAssignment{}
	err := elementLeader.FromDomainObject(v.ElementLeader)
	if err != nil {
		return errors.WithStack(err)
	}

	asstElementLeader := DutyAssignment{}
	err = asstElementLeader.FromDomainObject(v.AsstElementLeader)
	if err != nil {
		return errors.WithStack(err)
	}

	var members []ElementMember
	for _, m := range v.Members {
		member := ElementMember{ElementID: e.ID}
		err = member.FromDomainObject(m)
		if err != nil {
			return errors.WithStack(err)
		}

		members = append(members, member)
	}

	e.ID = v.ID()
	e.Name = v.Name
	e.ElementLeaderID = v.ElementLeader.ID()
	e.AsstElementLeaderID = v.AsstElementLeader.ID()
	e.ElementLeader = elementLeader
	e.AsstElementLeader = asstElementLeader
	e.Members = members

	return nil
}

func (e *Element) ToDomainObject() pkg.DomainObject {
	elementLeader := e.ElementLeader.ToDomainObject().(org.DutyAssignment)
	asstElementLeader := e.AsstElementLeader.ToDomainObject().(org.DutyAssignment)

	var members []org.Member
	for _, m := range e.Members {
		mbr := m.ToDomainObject().(org.Member)
		members = append(members, mbr)
	}

	return org.NewElement(e.ID, e.Name, elementLeader, asstElementLeader, members)
}

type ElementMember struct {
	ElementID uuid.UUID
	MemberID  uuid.UUID
	Member    Member
}

func (e *ElementMember) FromDomainObject(object pkg.DomainObject) error {
	if e.ElementID == uuid.Nil {
		return errors.New("failed to set ElementID on ElementMember before converting from domain object")
	}

	v, ok := object.(org.Member)
	if !ok {
		return errors.New("not a valid domain Member object")
	}

	e.MemberID = v.ID()

	return nil
}

func (e *ElementMember) ToDomainObject() pkg.DomainObject {
	return e.Member.ToDomainObject()
}
