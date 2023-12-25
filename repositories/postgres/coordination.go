package postgres

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

func (c *Coordination) Create() string {
	//TODO implement me
	panic("implement me")
}

func (c *Coordination) Fetch(eager bool) string {
	//TODO implement me
	panic("implement me")
}

func (c *Coordination) Update() string {
	//TODO implement me
	panic("implement me")
}

func (c *Coordination) UpdateOrCreate() string {
	//TODO implement me
	panic("implement me")
}

func (c *Coordination) Delete() string {
	//TODO implement me
	panic("implement me")
}

func (c *Coordination) Parameters(placeholder string, startIdx int) string {
	//TODO implement me
	panic("implement me")
}

func (c *Coordination) Values() []any {
	//TODO implement me
	panic("implement me")
}

func (c *Coordination) FromDomainObject(object pkg.DomainObject) error {
	panic(errors.New("Coordination.FromDomainObject() not implemented"))
}

func (c *Coordination) ToDomainObject() pkg.DomainObject {
	panic(errors.New("Coordination.ToDomainObject() not implemented"))
}
