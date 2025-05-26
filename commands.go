package main

import (
	"errors"
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
			name:        "explore {location name}",
			description: "Explore the Pokemons in the location area (e.g 'canalave-city-area')",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon name}",
			description: "Try to catch pokemon by their name (e.g 'wartortle')",
			callback:    commandCatch,
		},
		"list": {
			name:        "list",
			description: "List all caught Pokemons in Pokedex",
			callback:    commandList,
		},
		"inspect": {
			name:        "inspect {pokemon name}",
			description: "Inspect the properties of a Pokemon",
			callback:    commandInspect,
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
	locationAreaList, err := config.pokeApiClient.ListLocationAreas(config.nextPageURL)
	if err != nil {
		return err
	}

	config.nextPageURL = locationAreaList.Next
	config.previousPageURL = locationAreaList.Previous

	for _, locationArea := range locationAreaList.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMapBack(config *config, argument string) error {
	if config.previousPageURL == nil {
		fmt.Println("Can't move back. You are on first page!")
		return fmt.Errorf("unable to move back, config.PreviousPageURL: %v", config.previousPageURL)
	}
	locationAreaList, err := config.pokeApiClient.ListLocationAreas(config.previousPageURL)
	if err != nil {
		return err
	}

	config.nextPageURL = locationAreaList.Next
	config.previousPageURL = locationAreaList.Previous

	for _, locationArea := range locationAreaList.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandExplore(config *config, argument string) error {
	if argument == "" {
		return fmt.Errorf("location area name is missing")
	}
	locationArea, err := config.pokeApiClient.GetLocationArea(argument)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemons in %s:\n", locationArea.Name)
	for i, pokemonEncounter := range locationArea.PokemonEncounters {
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

	createTension()

	if tryCatch(pokemon.BaseExperience) {
		config.pokedex.Add(pokemon)
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You can now inspect it's details by using the 'inspect' command")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func commandList(config *config, argument string) error {
	pokedexPokemons := config.pokedex.List()
	if len(pokedexPokemons) == 0 {
		fmt.Println("No Pokemons in Pokedex")
		return nil
	}

	fmt.Println("Pokemons in Pokedex:")
	for i, pokemon := range pokedexPokemons {
		fmt.Printf("%d. %s\n", i+1, pokemon.Name)
	}
	return nil
}

func commandInspect(config *config, argument string) error {
	if argument == "" {
		return errors.New("pokemon name is missing")
	}
	pokemon, ok := config.pokedex.Entries[argument]
	if !ok {
		fmt.Printf("You haven't caught Pokemon: %s\n", argument)
		return nil
	}
	pokemon.Details()

	return nil
}
