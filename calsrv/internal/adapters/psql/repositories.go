package psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/models"
	uuid "github.com/satori/go.uuid"
)

func readRows(rows *sql.Rows) ([]models.Event, error) {
	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.ID, &event.Owner, &event.Title, &event.Text, &event.StartTime, &event.EndTime); err != nil {
			return []models.Event{}, fmt.Errorf("failed get events: %w", err)
		}
		events = append(events, event)
	}
	return events, nil
}

//AddEvent add new event
func (e *PostgresStorage) AddEvent(ctx context.Context, event models.Event) error {
	err := e.client.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	query := `INSERT INTO events(Owner_id, Title, Text, StartTime, EndTime)
				values ($1,$2,$3,$4,$5)`
	_, err = e.client.ExecContext(ctx, query,
		1, event.Title, event.Text, event.StartTime, event.EndTime)
	if err != nil {
		return fmt.Errorf("failed add event: %w", err)
	}
	return nil
}

//UpdateEvent Edit event
func (e *PostgresStorage) UpdateEvent(ctx context.Context, id uuid.UUID, newEvent models.Event) error {
	err := e.client.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	_, err = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update("events").
		Set("Owner_id", 1).
		Set("Title", newEvent.Title).
		Set("Text", newEvent.Text).
		Set("StartTime", newEvent.StartTime).
		Set("EndTime", newEvent.EndTime).
		Where(sq.Eq{"id": id.String()}).
		RunWith(e.client).ExecContext(ctx)

	if err != nil {
		return fmt.Errorf("failed update event: %w", err)
	}
	return nil
}

//GetEventByID Get event by ID
func (e *PostgresStorage) GetEventByID(ctx context.Context, id uuid.UUID) (models.Event, error) {
	nilResult := models.Event{}
	err := e.client.PingContext(ctx)
	if err != nil {
		return nilResult, fmt.Errorf("failed to connect: %w", err)
	}
	query := "SELECT id,owner_id,title,text,starttime,endtime from events WHERE id = $1"
	rows, err := e.client.QueryContext(ctx, query, id.String())
	if err != nil {
		return nilResult, fmt.Errorf("failed get event: %w", err)
	}
	defer rows.Close()
	var event models.Event
	if err = rows.Scan(&event.ID, &event.Owner, &event.Title, &event.Text, &event.StartTime, &event.EndTime); err != nil {
		return nilResult, fmt.Errorf("failed get event: %w", err)
	}
	return event, nil
}

//DelEvent Delete event
func (e *PostgresStorage) DelEvent(ctx context.Context, id uuid.UUID) error {
	err := e.client.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	query := `DELETE FROM events WHERE id = $1`
	_, err = e.client.ExecContext(ctx, query, id.String())
	if err != nil {
		return fmt.Errorf("failed delete event: %w", err)
	}
	return nil
}

//ListEvents Get slice of events
func (e *PostgresStorage) ListEvents(ctx context.Context, startdate time.Time, enddate time.Time) ([]models.Event, error) {
	nilResult := []models.Event{}
	err := e.client.PingContext(ctx)
	if err != nil {
		return nilResult, fmt.Errorf("failed to connect: %w", err)
	}
	fmt.Println(startdate, enddate)
	query := "SELECT id,owner_id,title,text,starttime,endtime FROM events WHERE starttime >= $1 and starttime <= $2 "
	rows, err := e.client.QueryContext(ctx, query, startdate, enddate)
	if err != nil {
		return nilResult, fmt.Errorf("failed get events: %w", err)
	}
	defer rows.Close()
	return readRows(rows)
}

//ListAllEvents Get slice of all events
func (e *PostgresStorage) ListAllEvents(ctx context.Context) ([]models.Event, error) {
	nilResult := []models.Event{}
	err := e.client.PingContext(ctx)
	if err != nil {
		return nilResult, fmt.Errorf("failed to connect: %w", err)
	}
	query := "SELECT * FROM events"
	rows, err := e.client.QueryContext(ctx, query)
	if err != nil {
		return nilResult, fmt.Errorf("failed get all events: %w", err)
	}
	defer rows.Close()
	return readRows(rows)
}
