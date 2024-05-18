package main

import (
	"bytes"
	"encoding/gob"
	"os"

	"github.com/timokae/pokedex/internal/pokeapi"
)

type state struct {
	NextLocationsUrl     string
	PreviousLocationsUrl string
	CaughtPokemon        map[string]pokeapi.Pokemon
}

func WriteStateToFile(cfg *config) error {
	stateToSave := NewStateFromConfig(cfg)

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	if err := encoder.Encode(stateToSave); err != nil {
		return err
	}

	file, err := os.Create("state")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func loadStateFromFile() (state, error) {
	input, err := os.ReadFile("state")
	if err != nil {
		return state{}, err
	}

	buffer := bytes.NewBuffer(input)
	decoder := gob.NewDecoder(buffer)

	loadedState := state{}

	if err := decoder.Decode(&loadedState); err != nil {
		return state{}, err
	}

	// fmt.Printf("LOADED %v\n", state.NextLocationsUrl)

	return loadedState, nil
}

func NewStateFromConfig(cfg *config) state {
	return state{
		NextLocationsUrl:     cfg.nextLocationsUrl,
		PreviousLocationsUrl: cfg.previousLocationsUrl,
		CaughtPokemon:        cfg.caughtPokemon,
	}
}
