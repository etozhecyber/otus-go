package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args, "<envdir> <command>")
		os.Exit(0)
	}

	//check envdir is exist
	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		log.Fatal("Error: envdir not exist - ", os.Args[1])
	}

	var env map[string]string
	env, _ = readDir(os.Args[1])
	os.Exit(runCmd(os.Args[2:], env))
}
