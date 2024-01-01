package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type DutyAssignment struct {
	ID           uuid.UUID
	OfficeSymbol string
	Title        string
	Assistant    bool
	DutyTitleID  uuid.UUID
	AssigneeID   *uuid.UUID
	DutyTitle    DutyTitle
	Assignee     *Member
}

func (da *DutyAssignment) FromDomainObject(object pkg.DomainObject) error {
	assignment, ok := object.(org.DutyAssignment)
	if !ok {
		return errors.New("attempt to pass non-org.DutyTitle object to DutyTitle.FromDomainObject")
	}

	dutyTitle := DutyTitle{}
	err := dutyTitle.FromDomainObject(assignment.DutyTitle())
	if err != nil {
		return errors.WithStack(err)
	}

	var assignee *Member
	if assignment.Assignee() != nil {
		assigneeID := assignment.Assignee().ID()
		da.AssigneeID = &assigneeID
		err = assignee.FromDomainObject(assignment.Assignee())
		if err != nil {
			return errors.WithStack(err)
		}
		da.Assignee = assignee
	}

	da.ID = assignment.ID()
	da.OfficeSymbol = assignment.OfficeSymbol()
	da.Title = assignment.Title()
	da.Assistant = assignment.Assistant()
	da.DutyTitleID = assignment.DutyTitle().ID()

	return nil
}

func (da *DutyAssignment) ToDomainObject() pkg.DomainObject {
	obj := da.DutyTitle.ToDomainObject()
	dutyTitle := obj.(org.DutyTitle)

	var assignee *org.Member
	if da.Assignee != nil {
		obj = da.Assignee.ToDomainObject()
		a := obj.(org.Member)
		assignee = &a
	}

	return org.NewDutyAssignment(
		da.ID,
		da.Title,
		da.OfficeSymbol,
		da.Assistant,
		dutyTitle,
		assignee,
	)
}

func (da *DutyAssignment) BeforeCreate(tx *gorm.DB) error {
	if da.ID != uuid.Nil {
		return nil
	}

	var err error
	da.ID, err = uuid.NewRandom()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
