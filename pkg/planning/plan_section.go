package planning

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type PlanSection struct {
	id    uuid.UUID
	Title string
	Body  string
}

func NewPlanSection(
	id uuid.UUID,
	title string,
	body string,
) PlanSection {
	return PlanSection{
		id:    id,
		Title: title,
		Body:  body,
	}
}

func (ps PlanSection) ID() uuid.UUID {
	return ps.id
}

func (ps PlanSection) LaTeX() string {
	panic(errors.New("PlanSection.LaTeX() not implemented"))
}
