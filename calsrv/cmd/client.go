package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var server, title, text, startTime, endTime string

const tsLayout = "2006-01-02T15:04:05"

//ClientCmd blabla
var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run grpc client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connect to: ", server)
	},
}

func init() {
	ClientCmd.Flags().StringVar(&server, "server", "localhost:8080", "host:port to connect to")
	ClientCmd.Flags().StringVar(&title, "title", "", "event title")
	ClientCmd.Flags().StringVar(&text, "text", "", "event text")
	ClientCmd.Flags().StringVar(&startTime, "start-time", "", "event start time, format: "+tsLayout)
	ClientCmd.Flags().StringVar(&endTime, "end-time", "", "event end time, format: "+tsLayout)
}
