package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	pokeCache map[string]cacheEntry
	mu        sync.Mutex
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	cacheMap := make(map[string]cacheEntry)
	cache := Cache{
		pokeCache: cacheMap,
	}
	go cache.reapLoop(interval)
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pokeCache[key] = cacheEntry{time.Now(), val}
} 

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.pokeCache[key]
	return entry.val, exists
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.reap(interval)
	}
}

func (c Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock() 
	for key, entry := range c.pokeCache {
		if time.Now().Sub(entry.createdAt) > interval {
			delete(c.pokeCache, key)
		}
	}
}