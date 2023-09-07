package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	commandOptions := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		if len(cleaned) == 0 {
			continue
		}

		command, ok := commandOptions[cleaned[0]]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(command string) []string {
	lowercase := strings.ToLower(command)
	words := strings.Split(lowercase, " ")
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCallback,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCallback,
		},
		"map": {
			name:        "map",
			description: "List locations",
			callback:    mapCallback,
		},
		"mapb": {
			name:        "mapb",
			description: "List locations (go back)",
			callback:    mapbCallback,
		},
		"explore": {
			name:        "explore (location area)",
			description: "Lists the Pokemon in a location area",
			callback:    exploreCallback,
		},
		"catch": {
			name:        "catch",
			description: "Try catching a Pokemon",
			callback:    catchCallback,
		},
		"inspect": {
			name:        "inspect",
			description: "Look at Pokemon that you've already caugh",
			callback:    inspectCallback,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List Pokemon in your Pokedex",
			callback:    pokedexCallback,
		},
	}
}
