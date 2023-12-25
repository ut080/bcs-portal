package postgres

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type DutyTitle struct {
	ID         uuid.UUID
	Title      string
	MemberType org.MemberType
	MinGrade   *org.Grade
	MaxGrade   *org.Grade
}

func (dt *DutyTitle) FromDomainObject(object pkg.DomainObject) error {
	v, ok := object.(org.DutyTitle)
	if !ok {
		if !ok {
			return errors.New("not a valid domain DutyTitle object")
		}
	}

	dt.ID = v.ID()
	dt.Title = v.Title
	dt.MemberType = v.MemberType
	dt.MinGrade = v.MinGrade
	dt.MaxGrade = v.MaxGrade

	return nil
}

func (dt *DutyTitle) ToDomainObject() pkg.DomainObject {
	return org.NewDutyTitle(dt.ID, dt.Title, dt.MemberType, dt.MinGrade, dt.MaxGrade)
}
