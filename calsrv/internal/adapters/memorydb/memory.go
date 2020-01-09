package memorydb

import (
	"context"
	"sync"

	"github.com/etozhecyber/otus-go/calsrv/internal/domain/models"
	uuid "github.com/satori/go.uuid"
)

//MemoryEventStorage struct
type MemoryEventStorage struct {
	mux sync.Mutex
	db  map[uuid.UUID]models.Event //simple index
}

//NewMemoryEventStorage create new memes
func NewMemoryEventStorage() (*MemoryEventStorage, error) {
	return &MemoryEventStorage{
		db: map[uuid.UUID]models.Event{},
	}, nil
}

//AddEvent add new event
func (m *MemoryEventStorage) AddEvent(ctx context.Context, event models.Event) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.db[event.ID] = event
	return nil
}

//GetEventByID Get event by ID
func (m *MemoryEventStorage) GetEventByID(ctx context.Context, id uuid.UUID) (models.Event, error) {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.db[id], nil
}

//DelEvent Delete event
func (m *MemoryEventStorage) DelEvent(ctx context.Context, id uuid.UUID) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	delete(m.db, id)
	return nil
}

//UpdateEvent Edit event
func (m *MemoryEventStorage) UpdateEvent(ctx context.Context, id uuid.UUID, newEvent models.Event) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.db[id] = newEvent
	return nil
}

//ListEvents Get slice of events
func (m *MemoryEventStorage) ListEvents(ctx context.Context) ([]models.Event, error) {
	var res []models.Event
	for _, event := range m.db {
		res = append(res, event)
	}
	return res, nil
}
