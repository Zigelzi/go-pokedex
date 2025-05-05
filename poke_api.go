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

func getLocationAreas(config *config, isMovingForwards bool) ([]locationArea, error) {
	endpoint := ""
	if isMovingForwards {
		endpoint = config.Next
	} else {
		endpoint = config.Previous
	}

	res, err := http.Get(endpoint)
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
