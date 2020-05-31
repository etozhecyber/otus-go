package main

import (
	"context"
	"errors"
	"fmt"
)

func worker(context context.Context, id int, jobs <-chan func() error, errs chan error) {
	for {
		select {
		case <-context.Done():
			println("worker", id, " quit")
			return
		case j, ok := <-jobs:
			if ok {
				err := j()
				errs <- err
			}
		}
	}
}

//Run jobs
func Run(task []func() error, numOfThreads int, maxErrors int) error {
	jobs := make(chan func() error, 100)
	errs := make(chan error, 1)
	ctx, cancel := context.WithCancel(context.Background())
	//creating workers
	for w := 1; w <= numOfThreads; w++ {
		go worker(ctx, w, jobs, errs)
	}
	//sending jobs for workers
	for _, j := range task {
		jobs <- j
	}
	close(jobs)

	// Control of working goroutines
	errnum := 0
	taskcomplite := 0
	for {
		select {
		//tracking errors
		case err := <-errs:
			if err != nil {
				errnum++
				if errnum >= maxErrors {
					cancel()
					return errors.New("Max errors")
				}
			}
			taskcomplite++
			if taskcomplite >= len(task) {
				cancel()
				return nil
			}

		}
	}
}
func main() {
	fmt.Println("use test")
}
