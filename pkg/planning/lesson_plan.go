package planning

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type LessonPlan struct {
	id         uuid.UUID
	Title      string
	Duration   uint
	Objectives string
	Resources  string
	Outline    string
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
		Title:      title,
		Duration:   duration,
		Objectives: objectives,
		Resources:  resources,
		Outline:    outline,
	}
}

func (lp LessonPlan) ID() uuid.UUID {
	return lp.id
}

func (lp LessonPlan) LaTeX() string {
	panic(errors.New("LessonPlan.LaTeX() not implemented"))
}
