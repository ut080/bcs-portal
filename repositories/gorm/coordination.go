package gorm

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/planning"
)

type Coordination struct {
	ID         uuid.UUID
	PlanID     uuid.UUID
	CoordOrder int
	OfficeID   uuid.UUID
	Action     planning.CoordinationAction
	Completed  time.Time
	Outcome    string
	Office     DutyAssignment
}

func (c *Coordination) FromDomainObject(object pkg.DomainObject) error {
	panic(errors.New("Coordination.FromDomainObject() not implemented"))
}

func (c *Coordination) ToDomainObject() pkg.DomainObject {
	panic(errors.New("Coordination.ToDomainObject() not implemented"))
}
