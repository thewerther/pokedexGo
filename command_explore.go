package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
  if len(args) < 1 {
    return errors.New("No area to explore given!")
  }

  locationDetailsResp, err := cfg.pokeClient.GetLocDetails(args[0])
  if err != nil {
    return err
  }

  fmt.Printf("Exploring %s...\n", locationDetailsResp.Name)
  for _, pokemon := range locationDetailsResp.PokemonEncounters {
    fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
  }

  return nil
}
