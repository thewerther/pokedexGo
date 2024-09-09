package pokeApi

import (
	"net/http"
	"time"
  "internal/pokeCache"
)

type PokeClient struct {
  cache       pokeCache.Cache
  httpClient  http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) PokeClient {
  return PokeClient{
    cache: pokeCache.NewCache(cacheInterval),
    httpClient: http.Client{
      Timeout: timeout,
    },
  }
}
