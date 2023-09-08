package airtable

import (
	"time"
)

type EntityFields interface {
	Key() string
}

type Entity struct {
	ID          string       `json:"id"`
	CreatedTime time.Time    `json:"createdTime"`
	Fields      EntityFields `json:"fields"`
}
