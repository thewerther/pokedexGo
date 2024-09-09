package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("No pokemon to inspect specified!")
	}

	pokemonName := args[0]
	if _, exists := cfg.caughtPokemon[pokemonName]; !exists {
		return errors.New(fmt.Sprintf("You have not caught %s yet!", pokemonName))
	}

	pokemonDetails, err := cfg.pokeClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf(`Name: %v
Height: %v
Weight: %v
Stats:
  - hp: %v
  - attack: %v
  - defense: %v
  - special-attack: %v
  - special-defense: %v
  - speed: %v
Types:
`,
  pokemonDetails.Name,
  pokemonDetails.Height,
  pokemonDetails.Weight,
  pokemonDetails.Stats[0].BaseStat,
  pokemonDetails.Stats[1].BaseStat,
  pokemonDetails.Stats[2].BaseStat,
  pokemonDetails.Stats[3].BaseStat,
  pokemonDetails.Stats[4].BaseStat,
  pokemonDetails.Stats[5].BaseStat,
	)
  for _, pokeType := range pokemonDetails.Types {
    fmt.Printf("  - %v\n", pokeType.Type.Name)
  }

	return nil
}
