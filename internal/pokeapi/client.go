package pokeapi

import (
	"net/http"
	"time"

	"github.com/Zigelzi/go-pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheLifetime time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(cacheLifetime),
	}
}
