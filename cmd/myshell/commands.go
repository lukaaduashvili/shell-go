package main

import (
	"fmt"
	"os"
)

type Command interface {
	Execute(args []string) error
}

type ExitCommand struct {
}

func (e *ExitCommand) Execute(args []string) error {
	os.Exit(0)
	return nil
}

type EchoCommand struct{}

func (e *EchoCommand) Execute(args []string) error {
	fmt.Println(args[0])
	return nil
}

const ()
