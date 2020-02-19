package services

import (
	"context"
	"testing"
	"time"

	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/memorydb"
	biserrors "github.com/etozhecyber/otus-go/calsrv/internal/domain/errors"
	"github.com/stretchr/testify/require"
)

func TestEventService(t *testing.T) {

	eventStorage, _ := memorydb.NewMemoryEventStorage()

	eventService := &EventService{
		EventStorage: eventStorage,
	}

	ctx := context.Background()

	startdate, _ := time.Parse("2006-01-02", "2020-01-01")
	enddate, _ := time.Parse("2006-01-02", "2020-01-10")
	weekenddate, _ := time.Parse("2006-01-02", "2020-01-04")

	//Creating event check
	require.Equal(t, nil, eventService.CreateEvent(ctx, "Andrey", "Drink water", "Drink cup of water", startdate, enddate))
	events, _ := eventService.GetAllEvents(ctx)
	require.Equal(t, len(events), 1)
	require.Equal(t, events[0].Owner, "Andrey")
	require.Equal(t, events[0].Title, "Drink water")
	require.Equal(t, events[0].Text, "Drink cup of water")
	require.Equal(t, events[0].StartTime, startdate)

	//Delete event check
	eventService.DelEventbyID(ctx, events[0].ID)
	events, _ = eventService.GetAllEvents(ctx)
	require.Equal(t, len(events), 0)

	/*****************
	 Bis errors checks
	*****************/

	//Check error date
	require.Equal(t, biserrors.EventError(biserrors.ErrIncorectEndDate), eventService.CreateEvent(ctx, "Andrey", "Drink water", "Drink cup of water", enddate, startdate))

	//Check error weekends errors
	require.Equal(t, biserrors.EventError(biserrors.ErrWeekendStartDate), eventService.CreateEvent(ctx, "Andrey", "Drink water", "Drink cup of water", weekenddate, enddate))
	require.Equal(t, biserrors.EventError(biserrors.ErrWeekendEndDate), eventService.CreateEvent(ctx, "Andrey", "Drink water", "Drink cup of water", startdate, weekenddate))
}
