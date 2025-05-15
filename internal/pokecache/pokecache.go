package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu         sync.RWMutex
	entries    map[string]cacheEntry
	lifetime   time.Duration
	totalBytes int
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(lifetime time.Duration) *Cache {
	newCache := &Cache{
		entries:  make(map[string]cacheEntry),
		lifetime: lifetime,
	}
	go newCache.reapLoop()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
	c.entries[key] = entry
	c.totalBytes += len(val)
	fmt.Printf("Cache size: %d KB \n", toKiloBytes(c.totalBytes))
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.lifetime)
	for {
		currentTime := <-ticker.C
		c.reap(currentTime)
	}
}

func (c *Cache) reap(currentTime time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entries {
		age := currentTime.Sub(entry.createdAt)
		if age >= c.lifetime {
			delete(c.entries, key)
		}
	}
}

func toKiloBytes(bytes int) int {
	const KB = 1024
	return bytes / KB
}
