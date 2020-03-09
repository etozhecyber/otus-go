package interfaces

import (
	"context"
	"time"

	"github.com/etozhecyber/otus-go/calsrv/internal/domain/models"
	uuid "github.com/satori/go.uuid"
)

/*
EventStorage interface of storage drivers
*/
type EventStorage interface {
	AddEvent(ctx context.Context, event models.Event) error
	GetEventByID(ctx context.Context, id uuid.UUID) (models.Event, error)
	DelEvent(ctx context.Context, id uuid.UUID) error
	UpdateEvent(ctx context.Context, id uuid.UUID, newEvent models.Event) error
	ListEvents(ctx context.Context, startdate time.Time, enddate time.Time) ([]models.Event, error)
	ListAllEvents(ctx context.Context) ([]models.Event, error)
}
