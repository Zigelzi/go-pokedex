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
	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area"
	}
	locationAreas, err := getLocationAreas(config, true)
	if err != nil {
		return err
	}
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMapBack(config *config) error {
	if config.Previous == "" {
		fmt.Println("Can't move back. You are on first page!")
		return fmt.Errorf("unable to move back, config.Previous: %s", config.Previous)
	}
	locationAreas, err := getLocationAreas(config, false)
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
			description: "Display next 20 location areas. Use multiple times to view more areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations areas. Use multiple times to view more areas",
			callback:    commandMapBack,
		},
	}
}
