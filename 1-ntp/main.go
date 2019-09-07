package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current time:", time.Format("02 January 2006 15:04:05.000"))
}
