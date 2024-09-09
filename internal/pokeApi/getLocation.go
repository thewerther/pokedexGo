package pokeApi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *PokeClient) ListLocations(pageURL *string) (LocationsResult, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

  if val, exists := c.cache.Get(url); exists {
    locationsResult := LocationsResult{}
    err := json.Unmarshal(val, &locationsResult)
    if err != nil {
      return LocationsResult{}, err
    }

    return locationsResult, nil
  }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsResult{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResult{}, err
	}
	defer resp.Body.Close()

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResult{}, err
	}

	locationsResp := LocationsResult{}
	err = json.Unmarshal(rawData, &locationsResp)
	if err != nil {
		return LocationsResult{}, err
	}

  c.cache.Add(url, rawData)
	return locationsResp, nil
}
