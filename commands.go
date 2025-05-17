package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
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
		"explore": {
			name:        "explore {location area name}",
			description: "Explore the Pokemons in the location area (e.g 'canalave-city-area')",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon name}",
			description: "Try to catch pokemon by their name (e.g 'wartortle')",
			callback:    commandCatch,
		},
	}
}

func commandExit(config *config, argument string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, argument string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println()

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMapForward(config *config, argument string) error {
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

func commandMapBack(config *config, argument string) error {
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

func commandExplore(config *config, argument string) error {
	if argument == "" {
		return fmt.Errorf("location area name is missing")
	}
	locationAreaResp, err := config.pokeApiClient.GetLocationAreaDetails(argument)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemons in %s:\n", locationAreaResp.Name)
	for i, pokemonEncounter := range locationAreaResp.PokemonEncounters {
		fmt.Printf("%d. %s\n", i+1, pokemonEncounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *config, argument string) error {
	if argument == "" {
		return fmt.Errorf("pokemon name is missing")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", argument)
	pokemon, err := config.pokeApiClient.GetPokemon(argument)
	if err != nil {
		return err
	}
	fmt.Printf("Found pokemon: %s\n", pokemon.Name)
	return nil
}
