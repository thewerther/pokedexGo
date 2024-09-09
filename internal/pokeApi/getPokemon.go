package pokeApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *PokeClient) GetPokemon(pokemonName string) (Pokemon, error) {
  url := baseURL + "/pokemon/" + pokemonName

  if val, exists := c.cache.Get(url); exists {
    pokemon := Pokemon{}
    err := json.Unmarshal(val, &pokemon)
    if err != nil {
      return Pokemon{}, err
    }

    return pokemon, nil
  }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
  }
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

  pokemonResp := Pokemon{}
  err = json.Unmarshal(rawData, &pokemonResp)
  if err != nil {
    if e, ok := err.(*json.SyntaxError); ok {
        fmt.Printf("syntax error at byte offset %d\n", e.Offset)
    }
    return Pokemon{}, err
  }

  c.cache.Add(url, rawData)

  return pokemonResp, nil
}
