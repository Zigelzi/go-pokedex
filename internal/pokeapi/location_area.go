package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListLocationAreas(pageURL *string) (ResponseLocationAreaList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, err := c.fetchData(url)
	if err != nil {
		return ResponseLocationAreaList{}, err
	}

	apiLocationAreaResponse := ResponseLocationAreaList{}
	err = json.Unmarshal(data, &apiLocationAreaResponse)
	if err != nil {
		return ResponseLocationAreaList{}, err
	}

	return apiLocationAreaResponse, nil
}

func (c *Client) GetLocationAreaDetails(locationName string) (ResponseLocationAreaDetails, error) {
	if locationName == "" {
		return ResponseLocationAreaDetails{}, fmt.Errorf("missing location name")
	}
	url := baseURL + "/location-area/" + locationName
	data, err := c.fetchData(url)
	if err != nil {
		return ResponseLocationAreaDetails{}, err
	}

	var locationAreaDetails ResponseLocationAreaDetails
	err = json.Unmarshal(data, &locationAreaDetails)
	if err != nil {
		return ResponseLocationAreaDetails{}, fmt.Errorf("unable to unmarshal LocationAreaDetails: %w", err)
	}
	return locationAreaDetails, nil
}
