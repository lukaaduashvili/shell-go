package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var commands = map[string]Command{
	"echo": &EchoCommand{},
	"exit": &ExitCommand{},
	"type": &TypeCommand{},
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
		fmt.Printf("%s is a shell builtin", queryCommand)
	} else {
		fmt.Printf("%s not found", queryCommand)
	}
	return nil
}
