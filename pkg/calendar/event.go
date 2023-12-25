package calendar

import (
	"time"

	"github.com/google/uuid"

	"github.com/ut080/bcs-portal/pkg/planning"
)

type Event struct {
	id    uuid.UUID
	title string
	start time.Time
	end   time.Time
	uod   *Uniform
	poc   string
	plan  *planning.Plan
}

func NewEvent(
	id uuid.UUID,
	title string,
	start time.Time,
	end time.Time,
	uod *Uniform,
	poc string,
	plan *planning.Plan,
) Event {
	return Event{
		id:    id,
		title: title,
		start: start,
		end:   end,
		uod:   uod,
		poc:   poc,
		plan:  plan,
	}
}

func (e Event) ID() uuid.UUID {
	return e.id
}

func (e Event) Title() string {
	return e.title
}

func (e Event) Start() time.Time {
	return e.start
}

func (e Event) End() time.Time {
	return e.end
}

func (e Event) UOD() *Uniform {
	return e.uod
}

func (e Event) POC() string {
	return e.poc
}

func (e Event) Plan() *planning.Plan {
	return e.plan
}
