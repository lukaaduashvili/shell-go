package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func requestInput(reader *bufio.Reader) string {
	fmt.Fprint(os.Stdout, "$ ")

	command, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	command = strings.TrimSpace(command)
	return command
}

func getCommandAndArgs(cliString string) (command string, args []string) {
	cliCommand := strings.Split(cliString, " ")
	command = cliCommand[0]
	args = cliCommand[1:]
	return command, args
}

func executeCommand(cliString string) {
	command, args := getCommandAndArgs(cliString)

	commandType, ok := commands[command]

	if !ok {
		err := executeExternalCommand(command, args)
		if err != nil {
			return
		}
	} else {
		err := commandType.Execute(args)
		if err != nil {
			return
		}
	}
}
