package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getLocationAreas(config *config) ([]locationArea, error) {
	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}
	res, err := http.Get(config.Next)
	if err != nil {
		return []locationArea{}, fmt.Errorf("unable to get locations: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []locationArea{}, err
	}
	locationResponse := locationResponse{}
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return []locationArea{}, err
	}

	config.Next = locationResponse.Next
	config.Previous = locationResponse.Previous

	return locationResponse.Results, nil
}
