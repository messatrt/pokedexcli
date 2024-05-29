package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("pokedex>")
		command, _ := reader.ReadString('\n')
		fmt.Printf(`enter %v
`, command)
	}
}
