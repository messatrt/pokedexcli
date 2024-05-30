package main

import "fmt"

func callBack_help(cfg *config) error {
	fmt.Println("Usage:")
	allCommands := getCommands()
	for _, command := range allCommands {
		fmt.Printf("-%s -%s\n", command.name, command.details)
	}
	return nil
}
