package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.Pokemon(args[0])
	if err != nil {
		return err
	}

	baseExp := float32(pokemon.BaseExperience)
	catchRate := 1.0 - (baseExp / 100.0 * rand.Float32())

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	if catchRate < 0.5 {
		fmt.Printf("%v escaped\n", pokemon.Name)
	} else {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	}

	return nil
}
