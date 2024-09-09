package pokeApi

import (
	"encoding/json"
	"io"
	"net/http"
  "fmt"
)

func (c *PokeClient) GetLocDetails(areaName string) (LocationDetails, error) {
  url := baseURL + "/location-area/" + areaName

  if val, exists := c.cache.Get(url); exists {
    locationDetailsResp := LocationDetails{}
    err := json.Unmarshal(val, &locationDetailsResp)
    if err != nil {
      return LocationDetails{}, err
    }

    return locationDetailsResp, nil
  }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
  }
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

  locationDetailsResp := LocationDetails{}
  err = json.Unmarshal(rawData, &locationDetailsResp)
  if err != nil {
    if e, ok := err.(*json.SyntaxError); ok {
        fmt.Printf("syntax error at byte offset %d\n", e.Offset)
    }
    return LocationDetails{}, err
  }

  c.cache.Add(url, rawData)
  return locationDetailsResp, nil
}
