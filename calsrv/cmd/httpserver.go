package cmd

import (
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//HTTPServerCmd server subprogram
var HTTPServerCmd = &cobra.Command{
	Use:   "httpserver",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utilities.GetConfiguration(configPath)
		if err != nil {
			log.Fatal(err)
		}
		err = utilities.SetupLogger(config)
		if err != nil {
			log.Fatal(err)
		}
		server, err := httpServerContruct()
		if err != nil {
			log.Fatal(err)
		}
		err = server.Serve(config)
		if err != nil {
			log.Fatal(err)
		}
	},
}
