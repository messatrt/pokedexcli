package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("pokedex> ")
		commandName, _ := reader.ReadString('\n')
		cleaned := cleanInput(commandName)
		if len(cleaned) == 0 {
			continue
		}
		commandName = cleaned[0]
		available := getCommands()
		command, ok := available[commandName]
		if !ok {
			fmt.Println("command not found")
			continue
		}
		command.callBack()
	}
}

type cliCommand struct {
	name     string
	details  string
	callBack func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			details:  "used to lookup help",
			callBack: callBack_help,
		},
		"exit": {
			name:     "exit",
			details:  "used to exit",
			callBack: callBack_exit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
