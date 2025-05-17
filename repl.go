package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Zigelzi/go-pokedex/internal/pokeapi"
	"github.com/Zigelzi/go-pokedex/internal/pokedex"
)

type config struct {
	pokeApiClient   *pokeapi.Client
	pokedex         *pokedex.Pokedex
	nextPageURL     *string
	previousPageURL *string
}

func startREPL(config *config) {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Type the command you want to do or write 'help' to view available commands")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Println("Please write the command that you want to perform")
			continue
		}

		commandName := words[0]
		commands := getCommands()
		command, ok := commands[commandName]
		if !ok {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		commandArgument := getCommandArgument(words)
		err := command.callback(config, commandArgument)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	text = strings.Trim(text, " ")
	words := strings.Split(text, " ")
	cleanedWords := make([]string, len(words))

	for i, word := range words {
		lowerCaseWord := strings.ToLower(word)
		cleanedWords[i] = lowerCaseWord
	}
	return cleanedWords
}

func getCommandArgument(input []string) string {
	if input == nil {
		return ""
	}
	if len(input) < 2 {
		return ""
	}
	return input[1]
}
