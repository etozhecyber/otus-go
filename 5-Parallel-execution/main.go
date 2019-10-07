package main

import (
	"errors"
	"fmt"
)

func parallelExecute(exec []func() error, numOfThreads int, maxErrors int) error {
	var start = make(chan int, numOfThreads)
	var errs = make(chan error)
	var done = make(chan string)

	//Creating goroutines
	for _, v := range exec {
		go func(f func() error) {
			start <- 1
			err := f()
			if err != nil {
				errs <- err
			}
			<-start
			done <- "done"
		}(v)
	}

	//Control of working goroutines
	errnum := 0
	donenum := 0
	for {
		select {
		//tracking errors
		case _ = <-errs:
			errnum++
			if errnum >= maxErrors {
				return errors.New("Max errors")
			}
		//Tracking of completed jobs
		case _ = <-done:
			donenum++
			if donenum >= len(exec) {
				return nil
			}
		}
	}
}

func main() {

	fmt.Println("use test")
}
