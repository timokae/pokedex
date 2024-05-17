package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) Pokemon(name string) (Pokemon, error) {
	if name == "" {
		return Pokemon{}, errors.New("please provide a pokemon name")
	}

	url := baseUrl + "pokemon/" + name
	body, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}

		if res.StatusCode == 404 {
			return Pokemon{}, errors.New("the pokemon could not be found")
		}

		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, body)
	}

	response := Pokemon{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return Pokemon{}, err
	}

	return response, nil
}
