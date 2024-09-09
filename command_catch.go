package main

import (
	"math/rand"
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
  if len(args) < 1 {
    return errors.New("No pokemon to catch specified!")
  }

  pokemonResp, err := cfg.pokeClient.GetPokemon(args[0])
  if err != nil {
    return err
  }

  fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)

  rNum := rand.Intn(pokemonResp.BaseExperience)
  if rNum > 50 {
    fmt.Printf("%s escaped!\n", pokemonResp.Name)
    return nil
  }

  fmt.Printf("%s was caught!\n", pokemonResp.Name)
  cfg.caughtPokemon[pokemonResp.Name] = pokemonResp

  return nil
}
