package planning

import (
	"github.com/google/uuid"

	"github.com/ut080/bcs-portal/pkg/org"
)

type TrainingBlock struct {
	id         uuid.UUID
	OPR        org.DutyAssignment
	CdtOPR     org.DutyAssignment
	Topic      string
	Instructor org.Member
	LessonPlan LessonPlan
}

func (tb TrainingBlock) ID() uuid.UUID {
	return tb.id
}
