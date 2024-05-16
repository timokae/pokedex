package main

import "fmt"

func commandHelp(config *config) error {
	commands := getCommands()

	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for key, command := range commands {
		fmt.Printf("%v: %v\n", key, command.description)
	}

	fmt.Println("")
	return nil
}
