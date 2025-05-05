package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println()

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(config *config) error {
	locationAreas, err := getLocationAreas(config)
	if err != nil {
		return err
	}
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Print the instructions of Pokedex",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display 20 location areas at once. Use multiple times to view more areas",
			callback:    commandMap,
		},
	}
}
