package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (ResponseLocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var data []byte

	data, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return ResponseLocationArea{}, fmt.Errorf("unable to get locations: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return ResponseLocationArea{}, err
		}
		c.cache.Add(url, data)
	}
	locationAreaResponse := ResponseLocationArea{}
	err := json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	return locationAreaResponse, nil
}
