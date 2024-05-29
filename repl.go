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
		command, _ := reader.ReadString('\n')
		cleaned := cleanInput(command)
		if len(cleaned) == 0 {
			continue
		}
		command = cleaned[0]
		switch command {
		case "exit":
			os.Exit(0)

		default:
			fmt.Println("command not found")
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
