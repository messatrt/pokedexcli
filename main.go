package main

import "github.com/messatrt/pokedexcli/internal/pokiapi"

type config struct {
	nextUrl       *string
	preUrl        *string
	pokiapiClient pokiapi.Client
}

func main() {
	cfg := config{
		pokiapiClient: pokiapi.NewClient(),
	}
	startRepl(&cfg)
}
