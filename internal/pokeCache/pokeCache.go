package pokeCache

import (
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
  createdAt time.Time
  val []byte
}

type Cache struct {
  entries map[string]CacheEntry
  mu *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
  newCache := Cache{
    entries: make(map[string]CacheEntry),
    mu: &sync.Mutex{},
  }
  go newCache.reapLoop(interval)

  return newCache
}

func (cache *Cache) reapLoop(interval time.Duration) {
  ticker := time.NewTicker(interval)
  for range ticker.C {
    cache.reap(time.Now().UTC(), interval)
  }
}

func (cache *Cache) reap(now time.Time, last time.Duration) {
  fmt.Println("Reaping cache entries")
  cache.mu.Lock()
  defer cache.mu.Unlock()
  for key, val := range cache.entries {
    if val.createdAt.Before(now.Add(-last)) {
      delete(cache.entries, key)
    }
  }
}

func (cache *Cache) Add(key string, val []byte) {
  cache.mu.Lock()
  defer cache.mu.Unlock()
  cache.entries[key] = CacheEntry{
    createdAt: time.Now().UTC(),
    val: val,
  }
}

func (cache *Cache) Get(key string) ([]byte, bool) {
  cache.mu.Lock()
  defer cache.mu.Unlock()
  foundVal, exists := cache.entries[key]
  if exists {
    fmt.Println("Cache hit for", key)
  } else {
    fmt.Println("Cache miss for", key)
  }

  return foundVal.val, exists
}
