package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func readDir(dir string) (env map[string]string, err error) {

	env = make(map[string]string)
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			env[info.Name()] = string(content)
		}
		return nil
	})
	return
}

func runCmd(cmd []string, env map[string]string) int {
	c1 := exec.Command(cmd[0], cmd[1:]...)
	c1.Env = append(os.Environ())
	for k, v := range env {
		c1.Env = append(c1.Env, k+"="+v)
	}

	c1.Stdout = os.Stdout
	c1.Stdin = os.Stdin
	c1.Stderr = os.Stderr
	c1.Start()
	if err := c1.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}
	return 0
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args, "<envdir> <command>")
		os.Exit(0)
	}

	//Test mode
	//Use like this "go run main.go envdir go run main.go test"
	if os.Args[1] == "test" {
		for _, v := range os.Environ() {
			fmt.Println(v)
		}
		os.Exit(99)
	}

	//check envdir is exist
	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		log.Fatal("Error: envdir not exist - ", os.Args[1])
	}

	var env map[string]string
	env, _ = readDir(os.Args[1])
	os.Exit(runCmd(os.Args[2:], env))
}
