package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/timokae/pokedex/internal/pokeapi"
)

type Command struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	nextLocationsUrl     *string
	previousLocationsUrl *string
	pokeapiClient        pokeapi.Client
	caughtPokemon        map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		command, exists := getCommands()[words[0]]
		args := words[1:]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of (the next) 20 locations in the Pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displayes the names of the previous 20 locations in the Pokemon World",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Lists all pokemon encounters in a area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Choose a pokemon by name and try to catch it!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "List details about a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all your caught pokemon",
			callback:    commandPokedex,
		},
	}
}
