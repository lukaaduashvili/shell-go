package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = map[string]string{
	"echo": ECHO,
}

func main() {
	runCLI()
}

func runCLI() {
	for {
		command := requestInput()
		executeCommand(command)
	}
}

func requestInput() string {
	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}
	command = strings.TrimSpace(command)
	return command
}

func executeCommand(command string) {
	_, ok := commands[command]

	if !ok {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}
