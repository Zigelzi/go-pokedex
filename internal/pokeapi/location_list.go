package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (ResponseLocationArea, error) {
	url := BASE_URL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return ResponseLocationArea{}, fmt.Errorf("unable to get locations: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocationArea{}, err
	}
	locationAreaResponse := ResponseLocationArea{}
	err = json.Unmarshal(body, &locationAreaResponse)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	return locationAreaResponse, nil
}
