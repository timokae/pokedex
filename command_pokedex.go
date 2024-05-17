package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your pokedex:")

	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %v\n", pokemon.Name)
	}

	return nil
}
