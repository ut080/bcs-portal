package postgres

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
)

type LessonPlan struct {
	ID         uuid.UUID
	Title      string
	Duration   uint
	Objectives string
	Resources  string
	Outline    string
}

func (lp *LessonPlan) FromDomainObject(object pkg.DomainObject) error {
	panic(errors.New("LessonPlan.FromDomainObject() not implemented"))
}

func (lp *LessonPlan) ToDomainObject() pkg.DomainObject {
	panic(errors.New("LessonPlan.ToDomainObject() not implemented"))
}
