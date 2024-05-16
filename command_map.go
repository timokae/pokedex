package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.LocationAreas(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationAreas.Next
	cfg.previousLocationsUrl = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousLocationsUrl == nil {
		return errors.New("you're on the first page")
	}

	locationAreas, err := cfg.pokeapiClient.LocationAreas(cfg.previousLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationAreas.Next
	cfg.previousLocationsUrl = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}