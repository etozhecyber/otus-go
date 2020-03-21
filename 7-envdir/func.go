package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func readDir(dir string) (env map[string]string, err error) {

	env = make(map[string]string)
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && !strings.ContainsRune(info.Name(), '=') {
			content, err := readFile(path)
			if err != nil {
				log.Println(err)
				return err
			}
			env[info.Name()] = string(content)
		}
		if info.Mode().IsDir() && info.Name() != dir {
			return filepath.SkipDir
		}
		return nil
	})
	return
}

func readFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	filescanner := bufio.NewScanner(file)
	filescanner.Scan()
	content := filescanner.Text()
	content = strings.TrimRight(content, " \t")
	content = strings.Replace(content, string(0x00), "\n", -1) //zero byte to new line conversion
	return string(content), err
}

func runCmd(cmd []string, env map[string]string) int {
	c1 := exec.Command(cmd[0], cmd[1:]...)

	// delete valiables
	for k, v := range env {
		if v == "" {
			os.Unsetenv(k)
			delete(env, k)
		}
	}

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
