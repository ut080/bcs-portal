package planning

import (
	"time"

	"github.com/google/uuid"
)

type MeetingPlan struct {
	id             uuid.UUID
	coordination   []Coordination
	PlanningStart  time.Time
	PlanDue        time.Time
	TrainingBlocks []TrainingBlock
}

func (mp *MeetingPlan) ID() uuid.UUID {
	return (*mp).id
}

func (mp *MeetingPlan) GetCoordination() []Coordination {
	return (*mp).coordination
}

func (mp *MeetingPlan) UpdateCoordination(idx int, coord Coordination) {
	if idx >= len((*mp).coordination) {
		(*mp).coordination = append((*mp).coordination, coord)
	} else {
		(*mp).coordination[idx] = coord
	}
}
