package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
)

type TrainingBlock struct {
	ID           uuid.UUID
	PlanID       uuid.UUID
	BlockNumber  int
	OPRID        uuid.UUID
	CdtOPRID     uuid.UUID
	Topic        string
	InstructorID uuid.UUID
	LessonPlanID uuid.UUID
	OPR          Member
	CdtOPR       Member
	LessonPlan   LessonPlan
}

func (tb *TrainingBlock) FromDomainObject(object pkg.DomainObject) error {
	panic(errors.New("TrainingBlock.FromDomainObject() not implemented"))
}

func (tb *TrainingBlock) ToDomainObject() pkg.DomainObject {
	panic(errors.New("TrainingBlock.ToDomainObject() not implemented"))
}
