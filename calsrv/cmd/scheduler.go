package cmd

import (
	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/scheduler"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

/*Scheduler hello command*/
var Scheduler = &cobra.Command{
	Use:   "scheduler",
	Short: "run notification sheduler",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utilities.GetConfiguration(configPath)
		if err != nil {
			log.Fatal(err)
		}
		err = utilities.SetupLogger(config)
		if err != nil {
			log.Fatal(err)
		}
		server, err := scheduler.NewNotifySheduler(config)
		if err != nil {
			log.Fatal(err)
		}
		err = server.Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}
