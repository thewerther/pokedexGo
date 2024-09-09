package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationsResp.Next
	cfg.prevURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeClient.ListLocations(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationResp.Next
	cfg.prevURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
