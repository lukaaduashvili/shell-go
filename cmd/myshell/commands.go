package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var commands = map[string]Command{
	"echo": &EchoCommand{},
	"exit": &ExitCommand{},
	"type": &TypeCommand{},
	"pwd":  &PwdCommand{},
	"cd":   &CdCommand{},
}

type Command interface {
	Execute(args []string) error
}

type ExitCommand struct {
}

func (e *ExitCommand) Execute(args []string) error {
	returnVal, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	os.Exit(returnVal)
	return nil
}

type EchoCommand struct{}

func (e *EchoCommand) Execute(args []string) error {
	fmt.Printf("%s\n", strings.Join(args, " "))
	return nil
}

type TypeCommand struct{}

func (t *TypeCommand) Execute(args []string) error {
	queryCommand := args[0]

	_, ok := commands[queryCommand]

	if ok {
		fmt.Printf("%s is a shell builtin\n", queryCommand)
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			fp := filepath.Join(path, args[0])
			if _, err := os.Stat(fp); err == nil {
				fmt.Println(fp)
				return nil
			}
		}
		fmt.Printf("%s not found\n", queryCommand)
	}
	return nil
}

type PwdCommand struct{}

func (p *PwdCommand) Execute(args []string) error {
	cmd := exec.Command("pwd", args...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	return err
}

type CdCommand struct{}

func (c *CdCommand) Execute(args []string) error {
	path := args[0]
	if path == "~" {
		path = os.Getenv("HOME")
	}
	err := os.Chdir(path)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", args[0])
	}
	return err
}

func executeExternalCommand(command string, args []string) error {
	if _, err := os.Stat(command); err == nil {
		cmd := exec.Command(command, args[:]...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()

		return err
	} else {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		return nil
	}
}
