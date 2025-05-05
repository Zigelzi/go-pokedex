package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config{
		Next:     "",
		Previous: "",
	}
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

		err := command.callback(&config)
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
