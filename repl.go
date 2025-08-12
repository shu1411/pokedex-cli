package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt = "Pokedex > "

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		if ok := scanner.Scan(); ok {
			userInput := cleanInput(scanner.Text())[0]

			if command, ok := commands[userInput]; !ok {
				fmt.Println("Unknown command.")
				continue
			} else {
				err := command.callback()
				fmt.Printf("%v", err)
			}
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
