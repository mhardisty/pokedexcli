package main

import "fmt"

func helpCallback(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are the available commands")

	availableCommands := getCommands()

	// find the length of the longest command name
	// allows for padding in 'whitespace_buffer' variable
	longest := 0
	for k := range availableCommands {
		if len(k) > longest {
			longest = len(k)
		}
	}

	// iterate over all commands and print them as name: description
	for k, v := range availableCommands {
		whitespace_buffer := ""

		for i := len(k); i < longest; i++ {
			whitespace_buffer += " "
		}

		fmt.Print(k, whitespace_buffer, ": ", v.description, "\n")

	}

	fmt.Println("")
	return nil
}
