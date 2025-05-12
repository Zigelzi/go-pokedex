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

	if cacheEntry, exists := c.cache.Get(url); exists {
		cachedLocationAreaResponse := ResponseLocationArea{}
		err := json.Unmarshal(cacheEntry, &cachedLocationAreaResponse)
		if err != nil {
			return ResponseLocationArea{}, err
		}
		return cachedLocationAreaResponse, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return ResponseLocationArea{}, fmt.Errorf("unable to get locations: %w", err)
	}
	defer res.Body.Close()

	apiData, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	apiLocationAreaResponse := ResponseLocationArea{}
	err = json.Unmarshal(apiData, &apiLocationAreaResponse)
	if err != nil {
		return ResponseLocationArea{}, err
	}

	c.cache.Add(url, apiData)
	return apiLocationAreaResponse, nil
}
