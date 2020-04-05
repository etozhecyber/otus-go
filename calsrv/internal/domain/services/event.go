package services

import (
	"context"
	"time"

	bisErrors "github.com/etozhecyber/otus-go/calsrv/internal/domain/errors"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/interfaces"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/models"
	uuid "github.com/satori/go.uuid"
)

// EventService struct
type EventService struct {
	EventStorage interfaces.EventStorage
}

// CreateEvent make new event
func (es *EventService) CreateEvent(ctx context.Context, owner int, title, text string, startTime, endTime time.Time) error {

	if startTime.After(endTime) {
		return bisErrors.EventError(bisErrors.ErrIncorectEndDate)
	}
	if startTime.Weekday() == 6 || startTime.Weekday() == 7 {
		return bisErrors.EventError(bisErrors.ErrWeekendStartDate)
	}
	if endTime.Weekday() == 6 || endTime.Weekday() == 7 {
		return bisErrors.EventError(bisErrors.ErrWeekendEndDate)
	}
	event := models.Event{
		ID:        uuid.NewV4(),
		Owner:     owner,
		Title:     title,
		Text:      text,
		StartTime: startTime,
		EndTime:   endTime,
	}
	err := es.EventStorage.AddEvent(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

// DelEventbyID delete event by ID
func (es *EventService) DelEventbyID(ctx context.Context, id uuid.UUID) error {
	err := es.EventStorage.DelEvent(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateEventbyID update event by ID
func (es *EventService) UpdateEventbyID(ctx context.Context, id uuid.UUID, event models.Event) error {
	err := es.EventStorage.UpdateEvent(ctx, id, event)
	if err != nil {
		return err
	}
	return nil
}

// GetEvents get events by date
func (es *EventService) GetEvents(ctx context.Context, startdate time.Time, enddate time.Time) ([]models.Event, error) {
	events, err := es.EventStorage.ListEvents(ctx, startdate, enddate)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetAllEvents get all events
func (es *EventService) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	events, err := es.EventStorage.ListAllEvents(ctx)
	if err != nil {
		return nil, err
	}
	return events, nil
}
