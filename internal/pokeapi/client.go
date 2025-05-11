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

func NewClient(timeout time.Duration) Client {
	const cacheLifetime = 10 * time.Second
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(cacheLifetime),
	}
}
