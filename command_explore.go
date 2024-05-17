package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a location name")
	}
	exploreRespone, err := cfg.pokeapiClient.Explore(&args[0])
	if err != nil {
		return err
	}

	for _, pokemon := range exploreRespone.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
