package main

import (
	"internal/pokeApi"
	"time"
)

func main() {
  pokeClient := pokeApi.NewClient(5 * time.Second, time.Minute * 5)
  cfg := &config{
    caughtPokemon: map[string]pokeApi.Pokemon{},
    pokeClient: pokeClient,
  }

  replLoop(cfg)
}
