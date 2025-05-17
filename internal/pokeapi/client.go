package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Zigelzi/go-pokedex/internal/pokecache"
)

type Client struct {
	httpClient  http.Client
	cache       pokecache.Cache
	cacheHits   int
	cacheMisses int
}

func NewClient(timeout, cacheLifetime time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(cacheLifetime),
	}
}

/*
fetchData tries to first get the data for the given URL from cache. If the URL doesn't exist in the cache
the data is fetched from PokeAPI directly
*/
func (c *Client) fetchData(url string) ([]byte, error) {
	if cacheEntry, exists := c.cache.Get(url); exists {
		c.cacheHits++
		// fmt.Printf("Cache hit for %s. Hit rate: %.1f\n", url, c.cacheHitRate())
		return cacheEntry, nil
	}
	c.cacheMisses++
	// fmt.Printf("Cache miss for %s. Hit rate: %.1f\n", url, c.cacheHitRate())

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to get data from URL: %s [%w]", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API responded with not OK: %s", res.Status)
	}
	apiData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, apiData)

	return apiData, nil
}

func (c *Client) cacheHitRate() float64 {
	total := c.cacheHits + c.cacheMisses
	if total == 0 {
		return 0
	}
	return float64(c.cacheHits) / float64(total) * 100
}
