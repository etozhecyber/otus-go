package main

import (
	"log"

	"github.com/etozhecyber/otus-go/calsrv/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
