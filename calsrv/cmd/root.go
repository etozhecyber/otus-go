package cmd

import (
	"github.com/spf13/cobra"
)

/*RootCmd default command*/
var RootCmd = &cobra.Command{
	Use:   "calsrv <server|client>",
	Short: "calsrv is a calendar server demo",
}

func init() {
	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(ClientCmd)
}
