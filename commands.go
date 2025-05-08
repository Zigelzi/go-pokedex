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

func commandMapForward(config *config) error {
	locationAreaResp, err := config.pokeApiClient.ListLocationAreas(config.nextPageURL)
	if err != nil {
		return err
	}

	config.nextPageURL = locationAreaResp.Next
	config.previousPageURL = locationAreaResp.Previous

	for _, locationArea := range locationAreaResp.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMapBack(config *config) error {
	if config.previousPageURL == nil {
		fmt.Println("Can't move back. You are on first page!")
		return fmt.Errorf("unable to move back, config.PreviousPageURL: %v", config.previousPageURL)
	}
	locationAreaResp, err := config.pokeApiClient.ListLocationAreas(config.previousPageURL)
	if err != nil {
		return err
	}

	config.nextPageURL = locationAreaResp.Next
	config.previousPageURL = locationAreaResp.Previous

	for _, locationArea := range locationAreaResp.Results {
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
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations areas. Use multiple times to view more areas",
			callback:    commandMapBack,
		},
	}
}
