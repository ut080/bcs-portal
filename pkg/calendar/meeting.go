package calendar

import (
	"github.com/google/uuid"
)

type Meeting struct {
	id    uuid.UUID
	Topic string
	Event Event
}

func NewMeeting(
	id uuid.UUID,
	topic string,
	event Event,
) Meeting {
	return Meeting{
		id:    id,
		Topic: topic,
		Event: event,
	}
}

func (m Meeting) ID() uuid.UUID {
	return m.id
}
