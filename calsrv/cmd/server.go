package cmd

import (
	"log"

	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/api"
	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/memorydb"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/services"

	"github.com/spf13/cobra"
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

var addr string
var dsn string

//ServerCmd server subprogram
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := construct()
		if err != nil {
			log.Fatal(err)
		}
		err = server.Serve(addr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	ServerCmd.Flags().StringVar(&addr, "addr", "localhost:8080", "host:port to listen")
	ServerCmd.Flags().StringVar(&dsn, "dsn", "host=127.0.0.1 user=event_user password=event_pwd dbname=event_db", "database connection string")
}
