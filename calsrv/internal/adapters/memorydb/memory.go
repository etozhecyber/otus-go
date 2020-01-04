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
	db  map[uuid.UUID]*models.Event //simple index
}

//NewMemoryEventStorage create new memes
func NewMemoryEventStorage() (*MemoryEventStorage, error) {
	return &MemoryEventStorage{
		db:  map[uuid.UUID]*models.Event{},
		mux: sync.Mutex{},
	}, nil
}

//AddEvent add new event
func (memes *MemoryEventStorage) AddEvent(ctx context.Context, event *models.Event) error {
	memes.mux.Lock()
	defer memes.mux.Unlock()
	memes.db[event.ID] = event
	return nil
}

//GetEventByID Get event by ID
func (memes *MemoryEventStorage) GetEventByID(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	return memes.db[id], nil
}

//DelEvent Delete event
func (memes *MemoryEventStorage) DelEvent(ctx context.Context, id uuid.UUID) error {
	memes.mux.Lock()
	defer memes.mux.Unlock()
	delete(memes.db, id)
	return nil
}

//UpdateEvent Edit event
func (memes *MemoryEventStorage) UpdateEvent(ctx context.Context, id uuid.UUID, newEvent *models.Event) error {
	memes.mux.Lock()
	defer memes.mux.Unlock()
	memes.db[id] = newEvent
	return nil
}

//ListEvents Get slice of events
func (memes *MemoryEventStorage) ListEvents(ctx context.Context) ([]*models.Event, error) {
	var res []*models.Event
	for _, event := range memes.db {
		res = append(res, event)
	}
	return res, nil
}
