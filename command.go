package main

import (
	"bufio"
	"fmt"
	"internal/pokeApi"
	"os"
	"strings"
)

type config struct {
  pokeClient      pokeApi.PokeClient
  caughtPokemon   map[string]pokeApi.Pokemon
  nextURL         *string
  prevURL         *string
}

type cliCommand struct {
  name        string
  description string
  callback    func(*config, ...string) error
  args        []string
}

func getAvailableCliCommands() map[string]cliCommand {
  return map[string]cliCommand {
    "help": {
      name:         "help",
      description:  "Displays a help message",
      callback:     commandHelp,
    },
    "exit": {
      name:         "exit",
      description:  "Exit the Pokedex",
      callback:     commandExit,
    },
    "map": {
      name:         "map",
      description:  "Prints the next 20 locations names.",
      callback:     commandMapf,
    },
    "mapb": {
      name:         "mapb",
      description:  "Prints the previous 20 location names.",
      callback:     commandMapb,
    },
    "explore": {
      name:         "explore",
      description:  "Explore an area by giving an area name.",
      callback:     commandExplore,
    },
    "catch": {
      name:         "catch",
      description:  "Try to catch a pokemon by specyfiying its name.",
      callback:     commandCatch,
    },
    "inspect": {
      name: "inspect",
      description: "Inspect a pokemon and print it's stats.",
      callback: commandInspect,
    },
    "pokedex": {
      name: "pokedex",
      description: "List all your caught pokemon.",
      callback: commandPokedex,
    },
  }
}

func sanitizeInput(inputText string) []string {
  lower := strings.ToLower(inputText)
  words := strings.Fields(lower)

  return words
}

func replLoop(cfg *config) {
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()

    inputWords := sanitizeInput(scanner.Text())
    if len(inputWords) == 0 {
      continue
    }

    cliCommand := inputWords[0]
    cliArgs := inputWords[1:]

    cmd, exists := getAvailableCliCommands()[cliCommand]
    if !exists {
      fmt.Printf("Command not found: %v\n", scanner.Text())
      continue
    }
    if err := scanner.Err(); err != nil {
      fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }

    err := cmd.callback(cfg, cliArgs...)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
  }
}
