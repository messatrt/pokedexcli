package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
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
		command.callBack(cfg)
	}
}

type cliCommand struct {
	name     string
	details  string
	callBack func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			details:  "used to lookup help",
			callBack: callBack_help,
		},
		"map": {
			name:     "map",
			details:  "shows the list of 20 locations",
			callBack: callBack_map,
		},
		"exit": {
			name:     "exit",
			details:  "used to exit",
			callBack: callBack_exit,
		},
		"map20": {
			name:     `map20`,
			details:  "show the prious map locaations",
			callBack: callBack_mapPre,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
