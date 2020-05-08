package main

import (
	"log"

	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/sender"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	"github.com/spf13/cobra"
)

//RootCmd default command
var RootCmd = &cobra.Command{
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

var configPath string

func init() {
	RootCmd.PersistentFlags().StringVar(&configPath, "config", "config.json", "path to config file")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
