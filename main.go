package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

var workingDir string

func main() {
	dir, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	} else {
		workingDir = path.Base(dir)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		_, _ = fmt.Fprintf(os.Stdout, "%s >> ", workingDir)
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		if err = execCmd(input); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execCmd(cmd string) error {
	//cmd = strings.TrimSuffix(cmd, "\n")
	cmd = strings.TrimSpace(cmd)
	parts := strings.Split(cmd, " ")

	switch parts[0] {
	case "cd":
		var err error
		var dir string
		if len(parts) < 2 {
			dir = os.Getenv("$HOME")
		} else {
			dir = parts[1]
		}
		if err = os.Chdir(dir); err == nil {
			curDir, err := os.Getwd()
			if err != nil {
				return err
			}
			workingDir = path.Base(curDir)
		}
		return err

	case "exit":
		os.Exit(0)
	}

	command := exec.Command(parts[0], parts[1:]...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	return command.Run()
}
