package pkg

import (
	"github.com/google/uuid"
)

type DomainObject interface {
	ID() uuid.UUID
}
