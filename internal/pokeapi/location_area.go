package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, err := c.fetchData(url)
	if err != nil {
		return LocationAreaList{}, err
	}

	locationAreaList := LocationAreaList{}
	err = json.Unmarshal(data, &locationAreaList)
	if err != nil {
		return LocationAreaList{}, err
	}

	return locationAreaList, nil
}

func (c *Client) GetLocationArea(locationName string) (LocationArea, error) {
	if locationName == "" {
		return LocationArea{}, fmt.Errorf("missing location name")
	}
	url := baseURL + "/location-area/" + locationName
	data, err := c.fetchData(url)
	if err != nil {
		return LocationArea{}, err
	}

	var locationArea LocationArea
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, fmt.Errorf("unable to unmarshal LocationArea: %w", err)
	}
	return locationArea, nil
}
