package planning

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CONPLAN struct {
	id           uuid.UUID
	coordination []Coordination
	planNumber   string
	title        string
	sections     []PlanSection
}

func NewCONPLAN(
	id uuid.UUID,
	coordination []Coordination,
	planNumber string,
	title string,
	sections []PlanSection,
) CONPLAN {
	return CONPLAN{
		id:           id,
		coordination: coordination,
		planNumber:   planNumber,
		title:        title,
		sections:     sections,
	}

}

func (c CONPLAN) ID() uuid.UUID {
	return c.id
}

func (c CONPLAN) Coordination() []Coordination {
	return c.coordination
}

func (c CONPLAN) PlanNumber() string {
	return c.planNumber
}

func (c CONPLAN) Title() string {
	return c.title
}

func (c CONPLAN) Sections() []PlanSection {
	return c.sections
}

func (c CONPLAN) LaTeX() string {
	panic(errors.New("CONPLAN.LaTeX() not implemented"))
}
