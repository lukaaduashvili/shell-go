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
	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}

	command = strings.TrimSpace(command)

	_, ok := commands[command]

	if !ok {
		fmt.Fprintf(os.Stdout, "%s: command not found", command)
	}

}
