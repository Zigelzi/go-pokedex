package pokedex

import (
	"fmt"
	"testing"

	"github.com/Zigelzi/go-pokedex/internal/pokeapi"
)

func TestAdd(t *testing.T) {
	zubat := pokeapi.Pokemon{Name: "zubat"}
	bulbasaur := pokeapi.Pokemon{Name: "bulbasaur"}
	venosaur := pokeapi.Pokemon{Name: "venosaur"}

	cases := []struct {
		name         string
		pokemons     []pokeapi.Pokemon
		expectedSize int
	}{
		{
			name:         "zero pokemons are added to pokedex",
			pokemons:     []pokeapi.Pokemon{},
			expectedSize: 0,
		},
		{
			name:         "all pokemons are added to pokedex",
			pokemons:     []pokeapi.Pokemon{zubat, bulbasaur, venosaur},
			expectedSize: 3,
		},
		{
			name:         "existing pokemons aren't added twice",
			pokemons:     []pokeapi.Pokemon{zubat, bulbasaur, bulbasaur},
			expectedSize: 2,
		},
	}

	for _, tc := range cases {
		pokedex := New()
		t.Run(tc.name, func(t *testing.T) {
			for _, pokemon := range tc.pokemons {
				pokedex.Add(pokemon)
			}

			fmt.Println(pokedex.Entries)

			if len(pokedex.Entries) != tc.expectedSize {
				t.Errorf("number of pokemons don't match: got [%d] want [%d]",
					len(pokedex.Entries),
					tc.expectedSize)
				return
			}

			for _, expectedPokemon := range tc.pokemons {
				_, exists := pokedex.Entries[expectedPokemon.Name]
				if !exists {
					t.Errorf("pokemon %s was not found in pokedex", expectedPokemon.Name)
					continue
				}
			}
		})
	}
}
