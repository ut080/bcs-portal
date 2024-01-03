package gorm_org

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type Unit struct {
	ID            uuid.UUID
	CharterNumber string
	UnitType      org.UnitType
	Name          string
	CommanderID   uuid.UUID
	Commander     DutyAssignment
}

func (u *Unit) FromDomainObject(object pkg.DomainObject) error {
	unit, ok := object.(org.Unit)
	if !ok {
		return errors.New("attempt to pass non-org.Unit object to Unit.FromDomainObject")
	}

	commander := DutyAssignment{}
	err := u.Commander.FromDomainObject(unit.Commander())
	if err != nil {
		return errors.WithStack(err)
	}

	u.ID = unit.ID()
	u.CharterNumber = unit.CharterNumber()
	u.UnitType = unit.UnitType()
	u.Name = unit.Name()
	u.CommanderID = unit.Commander().ID()
	u.Commander = commander

	return nil
}

func (u *Unit) ToDomainObject() pkg.DomainObject {
	obj := u.Commander.ToDomainObject()
	commander := obj.(org.DutyAssignment)

	return org.NewUnit(
		u.ID,
		u.CharterNumber,
		u.UnitType,
		u.Name,
		commander,
	)
}

func (u *Unit) BeforeCreate(tx *gorm.DB) error {
	if u.ID != uuid.Nil {
		return nil
	}

	var err error
	u.ID, err = uuid.NewRandom()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
