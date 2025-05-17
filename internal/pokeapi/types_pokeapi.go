package pokeapi

type ResponseLocationAreaList struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ResponseLocationAreaDetails struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	GameIndex         int                 `json:"game_index"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
}

type ResponsePokemonDetails struct {
	Pokemon Pokemon
}
