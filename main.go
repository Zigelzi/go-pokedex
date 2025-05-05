package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")
	for scan.Scan() {
		words := cleanInput(scan.Text())
		fmt.Printf("Your command was: %s\n", words[0])
		fmt.Print("Pokedex > ")
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
