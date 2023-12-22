package calendar

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	id    uuid.UUID
	Title string
	Start time.Time
	End   time.Time
	UOD   *Uniform
	POC   string
	//Plan *planning.Plan
}

func NewEvent(
	id uuid.UUID,
	title string,
	start time.Time,
	end time.Time,
	uod *Uniform,
	poc string,
	// Plan *planning.Plan
) Event {
	return Event{
		id:    id,
		Title: title,
		Start: start,
		End:   end,
		UOD:   uod,
		POC:   poc,
	}
}

func (e Event) ID() uuid.UUID {
	return e.id
}
