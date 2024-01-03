package gorm_org

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type DutyTitle struct {
	ID         uuid.UUID
	Code       string
	Title      string
	MemberType org.MemberType
	MinGrade   *org.Grade
	MaxGrade   *org.Grade
}

func (dt *DutyTitle) FromDomainObject(object pkg.DomainObject) error {
	title, ok := object.(org.DutyTitle)
	if !ok {
		return errors.New("attempt to pass non-org.DutyTitle object to DutyTitle.FromDomainObject")
	}

	dt.ID = title.ID()
	dt.Code = title.Code()
	dt.MemberType = title.MemberType()
	dt.MinGrade = title.MinGrade()
	dt.MaxGrade = title.MaxGrade()

	return nil
}

func (dt *DutyTitle) ToDomainObject() pkg.DomainObject {
	return org.NewDutyTitle(
		dt.ID,
		dt.Code,
		dt.Title,
		dt.MemberType,
		dt.MinGrade,
		dt.MaxGrade,
	)
}

func (dt *DutyTitle) BeforeCreate(tx *gorm.DB) error {
	if dt.ID != uuid.Nil {
		return nil
	}

	var err error
	dt.ID, err = uuid.NewRandom()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
