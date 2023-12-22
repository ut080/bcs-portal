package planning

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CONPLAN struct {
	id           uuid.UUID
	coordination []Coordination
	PlanNumber   string
	Title        string
	Sections     []PlanSection
}

func (c *CONPLAN) ID() uuid.UUID {
	return (*c).id
}

func (c *CONPLAN) GetCoordination() []Coordination {
	return (*c).coordination
}

func (c *CONPLAN) UpdateCoordination(idx int, coord Coordination) {
	if idx >= len((*c).coordination) {
		(*c).coordination = append((*c).coordination, coord)
	} else {
		(*c).coordination[idx] = coord
	}
}

func (c *CONPLAN) LaTeX() string {
	panic(errors.New("CONPLAN.LaTeX() not implemented"))
}
