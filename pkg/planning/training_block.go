package planning

import (
	"github.com/google/uuid"

	"github.com/ut080/bcs-portal/pkg/org"
)

type TrainingBlock struct {
	id         uuid.UUID
	opr        org.DutyAssignment
	cdtOPR     org.DutyAssignment
	topic      string
	instructor org.Member
	lessonPlan LessonPlan
}

func NewTrainingBlock(
	id uuid.UUID,
	opr org.DutyAssignment,
	cdtOPR org.DutyAssignment,
	topic string,
	instructor org.Member,
	lessonPlan LessonPlan,
) TrainingBlock {
	return TrainingBlock{
		id:         id,
		opr:        opr,
		cdtOPR:     cdtOPR,
		topic:      topic,
		instructor: instructor,
		lessonPlan: lessonPlan,
	}
}

func (tb TrainingBlock) ID() uuid.UUID {
	return tb.id
}

func (tb TrainingBlock) OPR() org.DutyAssignment {
	return tb.opr
}

func (tb TrainingBlock) CdtOPR() org.DutyAssignment {
	return tb.cdtOPR
}

func (tb TrainingBlock) Topic() string {
	return tb.topic
}

func (tb TrainingBlock) Instructor() org.Member {
	return tb.instructor
}

func (tb TrainingBlock) LessonPlan() LessonPlan {
	return tb.LessonPlan()
}
