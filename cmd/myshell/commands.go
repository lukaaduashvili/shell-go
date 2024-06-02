package main

import (
	"fmt"
	"os"
	"strconv"
)

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
	fmt.Println(args[0])
	return nil
}

const ()
