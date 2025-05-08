package main

import (
	"time"

	"github.com/Zigelzi/go-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeApiClient: pokeClient,
	}
	startREPL(config)
}
