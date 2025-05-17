package main

import (
	"time"

	"github.com/Zigelzi/go-pokedex/internal/pokeapi"
	"github.com/Zigelzi/go-pokedex/internal/pokedex"
)

func main() {
	const timeoutDuration = 5 * time.Second
	const cacheLifetime = 5 * time.Minute
	pokeClient := pokeapi.NewClient(timeoutDuration, cacheLifetime)
	config := &config{
		pokeApiClient: &pokeClient,
		pokedex:       pokedex.New(),
	}
	startREPL(config)
}
