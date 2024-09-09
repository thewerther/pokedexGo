package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
  fmt.Print(`
Welcome to the Pokedex!
Usage:
`)
  fmt.Println()
  for _, cmd := range getAvailableCliCommands() {
    fmt.Printf("\t%s: %s\n", cmd.name, cmd.description)
  }
  fmt.Println()

  return nil
}

