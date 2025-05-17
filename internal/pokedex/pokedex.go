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
