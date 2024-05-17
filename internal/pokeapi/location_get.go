package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationGet(area *string) (Location, error) {
	if area == nil {
		return Location{}, errors.New("area parameter missing")
	}

	url := baseUrl + "location-area/" + *area
	body, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return Location{}, err
		}

		if res.StatusCode == 404 {
			return Location{}, errors.New("the location could not be found")
		}

		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Location{}, err
		}

		c.cache.Add(url, body)
	}

	response := Location{}
	fmt.Println(url)
	err := json.Unmarshal(body, &response)
	if err != nil {
		return Location{}, err
	}

	return response, nil
}
