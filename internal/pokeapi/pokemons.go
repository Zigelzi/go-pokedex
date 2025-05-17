package pokeapi

import (
	"encoding/json"
	"errors"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	if name == "" {
		return Pokemon{}, errors.New("pokemon name is empty")
	}
	url := baseURL + "/pokemon/" + name
	data, err := c.fetchData(url)
	if err != nil {
		return Pokemon{}, err
	}
	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
