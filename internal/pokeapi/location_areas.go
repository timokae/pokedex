package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) LocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	url := baseUrl + "location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}

	body, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		c.cache.Add(url, body)
	}

	response := LocationAreasResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return response, nil
}
