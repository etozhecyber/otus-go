package main

import (
	"log"

	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/scheduler"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	"github.com/spf13/cobra"
)

//RootCmd default command
var RootCmd = &cobra.Command{
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

var configPath string

func init() {
	RootCmd.PersistentFlags().StringVar(&configPath, "config", "config.json", "path to config file")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
