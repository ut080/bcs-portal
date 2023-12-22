package planning

import (
	"time"

	"github.com/google/uuid"

	"github.com/ut080/bcs-portal/pkg/org"
)

type Coordination struct {
	id        uuid.UUID
	Office    org.DutyAssignment
	Action    CoordinationAction
	Completed time.Time
	Outcome   string
}

func NewCoordination(
	id uuid.UUID,
	office org.DutyAssignment,
	action CoordinationAction,
	completed time.Time,
	outcome string,
) Coordination {
	return Coordination{
		id:        id,
		Office:    office,
		Action:    action,
		Completed: completed,
		Outcome:   outcome,
	}
}

func (c Coordination) ID() uuid.UUID {
	return c.id
}
