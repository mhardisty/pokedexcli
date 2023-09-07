package main

import (
	"errors"
	"fmt"
)

func pokedexCallback(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("You don't have any pokemon yet! Try the catch command!")
	}
	fmt.Println("Your Pokemon")
	fmt.Println("------------")
	for k := range cfg.caughtPokemon {
		fmt.Println(k)
	}
	return nil
}
