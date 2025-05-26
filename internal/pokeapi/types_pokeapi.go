package pokeapi

import "fmt"

type LocationAreaList struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	ID                int                 `json:"id"`
	URL               string              `json:"url"`
	Name              string              `json:"name"`
	GameIndex         int                 `json:"game_index"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	URL            string        `json:"url"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

func (p *Pokemon) Details() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d cm\n", p.Height*10)
	fmt.Printf("Weight: %d kg\n", p.Weight/10)
	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStatValue)
	}

	fmt.Println("Types:")
	for _, pokemonType := range p.Types {
		fmt.Printf("  -%s\n", pokemonType.Type.Name)
	}
}

type PokemonStat struct {
	Stat          Stat `json:"stat"`
	BaseStatValue int  `json:"base_stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type Type `json:"type"`
}
type Type struct {
	Name string `json:"name"`
}
