package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

// CtrlD is CTRL + D keycode
const CtrlD byte = 4

func main() {

	var timeout string
	flag.StringVar(&timeout, "timeout", "60s", "timeout operation")
	flag.Parse()
	if flag.Arg(1) == "" {
		log.Fatalln("Need enter host and port")
	}
	duration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalln("Parse duration error:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)

	//this code need for raw read/write from/to terminal
	state, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatalln("setting stdin to raw:", err)
	}
	defer func() {
		log.Println("restore terminal")
		if err := terminal.Restore(0, state); err != nil {
			log.Println("warning, failed to restore terminal:", err)
		}
	}()

	//Create connection
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", flag.Arg(0)+":"+flag.Arg(1))
	if err != nil {
		log.Fatalln("Cannot connect:", err)
	}

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		netRead(ctx, conn)
		log.Println("NetRead exit")
		cancel()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		netSend(ctx, conn)
		log.Println("NetSend exit")
		cancel()
		wg.Done()
	}()

	wg.Wait()
	conn.Close()
}

func netRead(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanBytes)
OUT:
	for {
		select {
		case <-ctx.Done():
			break OUT
		default:
			if !scanner.Scan() {
				break OUT
			}
			str := scanner.Text()
			fmt.Print(str)
		}
	}
}

func netSend(ctx context.Context, conn net.Conn) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
OUT:
	for {
		select {
		case <-ctx.Done():
			break OUT
		default:
			if !scanner.Scan() {
				log.Println("error scanner scan")
				break OUT
			}
			// IF CTRL+D - exit
			if scanner.Bytes()[0] == CtrlD {
				break OUT
			}
			fmt.Print(scanner.Text())
			conn.Write(scanner.Bytes())
		}
	}
}
