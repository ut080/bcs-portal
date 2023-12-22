package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

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
	v, ok := object.(org.DutyAssignment)
	if !ok {
		return errors.New("not a valid domain DutyAssignment object")
	}

	dutyTitle := DutyTitle{}
	err := dutyTitle.FromDomainObject(v.DutyTitle)
	if err != nil {
		return errors.WithStack(err)
	}

	if v.Assignee != nil {
		assigneeID := v.Assignee.ID()
		da.AssigneeID = &assigneeID

		assignee := Member{}
		err = assignee.FromDomainObject(*v.Assignee)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	da.ID = v.ID()
	da.OfficeSymbol = v.OfficeSymbol
	da.Title = v.Title
	da.Assistant = v.Assistant
	da.DutyTitleID = v.DutyTitle.ID()

	return nil
}

func (da *DutyAssignment) ToDomainObject() pkg.DomainObject {
	dutyTitle := da.DutyTitle.ToDomainObject().(org.DutyTitle)

	var assignee *org.Member
	if da.Assignee != nil {
		a := da.Assignee.ToDomainObject().(org.Member)
		assignee = &a
	}

	return org.NewDutyAssignment(da.ID, da.Title, da.OfficeSymbol, da.Assistant, dutyTitle, assignee)
}
