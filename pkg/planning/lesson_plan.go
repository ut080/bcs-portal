package planning

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type LessonPlan struct {
	id         uuid.UUID
	title      string
	duration   uint
	objectives string
	resources  string
	outline    string
}

func NewLessonPlan(
	id uuid.UUID,
	title string,
	duration uint,
	objectives string,
	resources string,
	outline string,
) LessonPlan {
	return LessonPlan{
		id:         id,
		title:      title,
		duration:   duration,
		objectives: objectives,
		resources:  resources,
		outline:    outline,
	}
}

func (lp LessonPlan) ID() uuid.UUID {
	return lp.id
}

func (lp LessonPlan) Title() string {
	return lp.title
}

func (lp LessonPlan) Duration() uint {
	return lp.duration
}

func (lp LessonPlan) Objectives() string {
	return lp.objectives
}

func (lp LessonPlan) Resources() string {
	return lp.resources
}

func (lp LessonPlan) Outline() string {
	return lp.outline
}

func (lp LessonPlan) LaTeX() string {
	panic(errors.New("LessonPlan.LaTeX() not implemented"))
}
