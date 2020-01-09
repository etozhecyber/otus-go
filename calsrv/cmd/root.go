package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd default command
var RootCmd = &cobra.Command{
	Use:   "calsrv <server|client>",
	Short: "calsrv is a calendar server demo",
}

var configPath string

func init() {
	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(ClientCmd)
	RootCmd.AddCommand(HelloCmd)
	RootCmd.PersistentFlags().StringVar(&configPath, "config", "config.json", "path to config file")
}
