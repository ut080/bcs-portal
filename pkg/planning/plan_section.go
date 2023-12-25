package planning

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type PlanSection struct {
	id    uuid.UUID
	title string
	body  string
}

func NewPlanSection(
	id uuid.UUID,
	title string,
	body string,
) PlanSection {
	return PlanSection{
		id:    id,
		title: title,
		body:  body,
	}
}

func (ps PlanSection) ID() uuid.UUID {
	return ps.id
}

func (ps PlanSection) Title() string {
	return ps.title
}

func (ps PlanSection) Body() string {
	return ps.body
}

func (ps PlanSection) LaTeX() string {
	panic(errors.New("PlanSection.LaTeX() not implemented"))
}
