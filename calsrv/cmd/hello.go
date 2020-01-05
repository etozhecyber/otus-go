package cmd

import (
	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/hello"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

/*HelloCmd hello command*/
var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Run hello server",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utilities.GetConfiguration(configPath)
		if err != nil {
			log.Fatal(err)
		}
		err = utilities.SetupLogger(config)
		if err != nil {
			log.Fatal(err)
		}
		err = hello.Serve(config)
		if err != nil {
			log.Fatal(err)
		}
	},
}
