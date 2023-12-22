package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/calendar"
)

type Meeting struct {
	ID      uuid.UUID
	EventID uuid.UUID
	Topic   string
	Event   Event
}

func (m *Meeting) FromDomainObject(object pkg.DomainObject) error {
	meeting, ok := object.(calendar.Meeting)
	if !ok {
		return errors.New("not a valid domain [OBJECT] object")
	}

	event := Event{}
	err := event.FromDomainObject(meeting.Event)
	if err != nil {
		return errors.WithStack(err)
	}

	m.ID = meeting.ID()
	m.EventID = meeting.Event.ID()
	m.Topic = meeting.Topic
	m.Event = event

	return nil
}

func (m *Meeting) ToDomainObject() pkg.DomainObject {
	event := m.Event.ToDomainObject().(calendar.Event)

	return calendar.NewMeeting(m.ID, m.Topic, event)
}
