package main

import (
	"time"

	"github.com/Zigelzi/go-pokedex/internal/pokeapi"
)

func main() {
	const timeoutDuration = 5
	pokeClient := pokeapi.NewClient(timeoutDuration * time.Second)
	config := &config{
		pokeApiClient: &pokeClient,
	}
	startREPL(config)
}
