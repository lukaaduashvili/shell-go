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
	reader := bufio.NewReader(os.Stdin)
	for {
		command := requestInput(reader)
		executeCommand(command)
	}
}

func requestInput(reader bufio.Reader) string {
	fmt.Fprint(os.Stdout, "$ ")

	command, err := reader.ReadString('\n')
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
