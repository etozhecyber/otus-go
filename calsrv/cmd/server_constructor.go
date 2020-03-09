package cmd

import (
	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/api"
	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/httpapi"
	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/memorydb"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/services"
)

// TODO: dependency injection, orchestrator
func construct() (*api.CalendarServer, error) {
	eventStorage, err := memorydb.NewMemoryEventStorage()
	if err != nil {
		return nil, err
	}
	eventService := &services.EventService{
		EventStorage: eventStorage,
	}
	server := &api.CalendarServer{
		EventUsecases: eventService,
	}
	return server, nil
}

func httpServerContruct() (*httpapi.HTTPServer, error) {
	eventStorage, err := memorydb.NewMemoryEventStorage()
	if err != nil {
		return nil, err
	}
	eventService := &services.EventService{
		EventStorage: eventStorage,
	}
	server := &httpapi.HTTPServer{
		EventUsecases: eventService,
	}
	return server, nil
}
