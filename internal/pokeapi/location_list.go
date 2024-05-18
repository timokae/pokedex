package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) LocationList(pageUrl string) (ShallowLocation, error) {
	url := baseUrl + "location-area"
	if pageUrl != "" {
		url = pageUrl
	}

	body, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return ShallowLocation{}, err
		}

		if res.StatusCode == 404 {
			return ShallowLocation{}, errors.New("the location could not be found")
		}

		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return ShallowLocation{}, err
		}

		c.cache.Add(url, body)
	}

	response := ShallowLocation{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return ShallowLocation{}, err
	}

	return response, nil
}
