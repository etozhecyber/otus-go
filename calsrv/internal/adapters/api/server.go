package api

import (
	"fmt"
	"time"

	"github.com/etozhecyber/otus-go/calsrv/internal/domain/services"
)

//CalendarServer stub
type CalendarServer struct {
	EventUsecases *services.EventService
}

//Serve stub
func (cs *CalendarServer) Serve(addr string) error {
	fmt.Println("start server on:", addr)
	time.Sleep(2 * time.Minute)
	return nil
}
