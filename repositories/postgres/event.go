package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/calendar"
)

type Event struct {
	ID            uuid.UUID
	Title         string
	StartDatetime time.Time
	EndDatetime   time.Time
	UOD           *calendar.Uniform
	POC           string
}

func (e *Event) FromDomainObject(object pkg.DomainObject) error {
	event, ok := object.(calendar.Event)
	if !ok {
		return errors.New("not a valid domain [OBJECT] object")
	}

	e.ID = event.ID()
	e.Title = event.Title
	e.StartDatetime = event.Start
	e.EndDatetime = event.End
	e.UOD = event.UOD
	e.POC = event.POC

	return nil
}

func (e *Event) ToDomainObject() pkg.DomainObject {
	return calendar.NewEvent(e.ID, e.Title, e.StartDatetime, e.EndDatetime, e.UOD, e.POC)
}
