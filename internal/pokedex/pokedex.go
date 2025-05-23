package pokedex

import (
	"sync"

	"github.com/Zigelzi/go-pokedex/internal/pokeapi"
)

type Pokedex struct {
	Entries map[string]pokeapi.Pokemon
	mu      sync.RWMutex
}

func New() *Pokedex {
	return &Pokedex{
		Entries: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Entries[pokemon.Name] = pokemon
}

// List all Pokemon in the Pokedex that user has caught.
func (p *Pokedex) List() []pokeapi.Pokemon {
	p.mu.RLock()
	defer p.mu.RUnlock()

	var pokedexEntries []pokeapi.Pokemon
	for _, pokemon := range p.Entries {
		pokedexEntries = append(pokedexEntries, pokemon)
	}
	return pokedexEntries
}
