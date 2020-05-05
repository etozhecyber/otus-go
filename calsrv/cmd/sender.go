package cmd

import (
	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/sender"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

/*Sender hello command*/
var Sender = &cobra.Command{
	Use:   "sender",
	Short: "run notification sender",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utilities.GetConfiguration(configPath)
		if err != nil {
			log.Fatal(err)
		}
		err = utilities.SetupLogger(config)
		if err != nil {
			log.Fatal(err)
		}
		server, err := sender.NewNotifySender(config)
		if err != nil {
			log.Fatal(err)
		}
		err = server.Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}
