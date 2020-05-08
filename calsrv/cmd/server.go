package cmd

import (
	"log"

	"github.com/etozhecyber/otus-go/calsrv/utilities"
	"github.com/spf13/cobra"
)

//ServerCmd server subprogram
var ServerCmd = &cobra.Command{
	Use:   "server",
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
