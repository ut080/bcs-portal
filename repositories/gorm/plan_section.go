package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
)

type PlanSection struct {
	ID            uuid.UUID
	PlanID        uuid.UUID
	SectionNumber int
	Title         string
	Body          string
}

func (ps *PlanSection) FromDomainObject(object pkg.DomainObject) error {
	panic(errors.New("PlanSection.FromDomainObject() not implemented"))
}

func (ps *PlanSection) ToDomainObject() pkg.DomainObject {
	panic(errors.New("PlanSection.ToDomainObject() not implemented"))
}
