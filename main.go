package main

import (
	"time"

	"github.com/Zigelzi/go-pokedex/internal/pokeapi"
)

func main() {
	const timeoutDuration = 5 * time.Second
	const cacheLifetime = 5 * time.Minute
	pokeClient := pokeapi.NewClient(timeoutDuration, cacheLifetime)
	config := &config{
		pokeApiClient: &pokeClient,
	}
	startREPL(config)
}
