package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
  if len(cfg.caughtPokemon) == 0 {
    return errors.New("You have not caught any pokemon yet!")
  }

  fmt.Println("Your pokedex:")
  for _, pokemon := range cfg.caughtPokemon {
    fmt.Printf("  - %v\n", pokemon.Name)
  }

  return nil
}
