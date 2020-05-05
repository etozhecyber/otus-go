package main

import (
	"log"

	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/psql"
	grpc "github.com/etozhecyber/otus-go/calsrv/internal/adapters/server"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/services"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	"github.com/spf13/cobra"
)

func grpcServerConstruct(config utilities.Config) (*grpc.CalendarServer, error) {
	eventStorage, err := psql.NewPostgresStorage(config)
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

//RootCmd default command
var RootCmd = &cobra.Command{
	Use:   "api server",
	Short: "Run GRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utilities.GetConfiguration(configPath)
		if err != nil {
			log.Fatal(err)
		}
		err = utilities.SetupLogger(config)
		if err != nil {
			log.Fatal(err)
		}
		server, err := grpcServerConstruct(config)
		if err != nil {
			log.Fatal(err)
		}
		err = server.Serve(config)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var configPath string

func init() {
	RootCmd.PersistentFlags().StringVar(&configPath, "config", "config.json", "path to config file")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
