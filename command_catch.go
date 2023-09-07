package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catchCallback(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	const THRESHOLD = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > THRESHOLD {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}
	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("You caught %s!\n", pokemonName)
	return nil
}
