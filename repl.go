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
		fmt.Print("pokedex>")
		command, _ := reader.ReadString('\n')
		cleaned := cleanInput(command)
		fmt.Printf(`enter %v
`, cleaned)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
