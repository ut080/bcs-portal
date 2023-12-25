package planning

import (
	"time"

	"github.com/google/uuid"
)

type MeetingPlan struct {
	id             uuid.UUID
	coordination   []Coordination
	planningStart  time.Time
	planDue        time.Time
	trainingBlocks []TrainingBlock
}

func NewMeetingPlan(
	id uuid.UUID,
	coordination []Coordination,
	planningStart time.Time,
	planDue time.Time,
	trainingBlocks []TrainingBlock,
) MeetingPlan {
	return MeetingPlan{
		id:             id,
		coordination:   coordination,
		planningStart:  planningStart,
		planDue:        planDue,
		trainingBlocks: trainingBlocks,
	}
}

func (mp MeetingPlan) ID() uuid.UUID {
	return mp.id
}

func (mp MeetingPlan) Coordination() []Coordination {
	return mp.coordination
}

func (mp MeetingPlan) PlanningStart() time.Time {
	return mp.planningStart
}

func (mp MeetingPlan) PlanDue() time.Time {
	return mp.planDue
}

func (mp MeetingPlan) TrainingBlocks() []TrainingBlock {
	return mp.trainingBlocks
}
