package calendar

import (
	"github.com/google/uuid"
)

type Meeting struct {
	id    uuid.UUID
	topic string
	event Event
}

func NewMeeting(
	id uuid.UUID,
	topic string,
	event Event,
) Meeting {
	return Meeting{
		id:    id,
		topic: topic,
		event: event,
	}
}

func (m Meeting) ID() uuid.UUID {
	return m.id
}

func (m Meeting) Topic() string {
	return m.topic
}

func (m Meeting) Event() Event {
	return m.event
}
