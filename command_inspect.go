package main

import (
	"errors"
	"fmt"
)

func inspectCallback(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Which Pokemon would you like to inspect?")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("You haven't caught that Pokemon yet...")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Println("----------------------------------")
	for i, kind := range pokemon.Types {
		fmt.Printf("Type %v: %v\n", i+1, kind.Type.Name)

	}
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("  %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
