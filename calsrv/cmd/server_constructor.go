package cmd

import (
	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/httpapi"
	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/memorydb"
	grpc "github.com/etozhecyber/otus-go/calsrv/internal/adapters/server"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/services"
)

func grpcServerConstruct() (*grpc.CalendarServer, error) {
	eventStorage, err := memorydb.NewMemoryEventStorage()
	if err != nil {
		return nil, err
	}
	eventService := &services.EventService{
		EventStorage: eventStorage,
	}
	server := &grpc.CalendarServer{
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
