package pokeapi

import (
	"net/http"
	"time"

	"github.com/havokmoobii/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient   http.Client
	pokeapiCache pokecache.Cache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	cache := pokecache.NewCache(cacheInterval)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeapiCache: cache,
	}
}