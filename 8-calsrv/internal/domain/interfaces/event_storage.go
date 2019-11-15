package interfaces

import (
	"context"

	"github.com/etozhecyber/otus-go/8-calsrv/internal/domain/models"
)

/*
EventStorage interface
*/
type EventStorage interface {
	AddEvent(ctx context.Context, event *models.Event) error
	DelEvent(ctx context.Context, event *models.Event) error
	EditEvent(ctx context.Context, oldEvent *models.Event, newEvent *models.Event)
	GetEventById(ctx context.Context, id string) (*models.Event, error)
}
