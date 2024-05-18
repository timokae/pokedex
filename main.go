package main

import (
	"fmt"
	"time"

	"github.com/timokae/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	loadedState, err := loadStateFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg := &config{
		pokeapiClient:        pokeClient,
		nextLocationsUrl:     loadedState.NextLocationsUrl,
		previousLocationsUrl: loadedState.PreviousLocationsUrl,
		caughtPokemon:        loadedState.CaughtPokemon,
	}

	startRepl(cfg)
}
