package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Event model
type Event struct {
	ID        uuid.UUID
	Owner     string
	Title     string
	Text      string
	StartTime time.Time
	EndTime   time.Time
}
