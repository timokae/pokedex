package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Explore(area *string) (ExploreRespone, error) {
	if area == nil {
		return ExploreRespone{}, errors.New("area parameter missing")
	}

	url := baseUrl + "location-area/" + *area
	body, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return ExploreRespone{}, err
		}

		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return ExploreRespone{}, err
		}

		c.cache.Add(url, body)
	}

	response := ExploreRespone{}
	fmt.Println(url)
	err := json.Unmarshal(body, &response)
	if err != nil {
		return ExploreRespone{}, err
	}

	return response, nil
}
