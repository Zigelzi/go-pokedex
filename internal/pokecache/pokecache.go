package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	Entry    map[string]cacheEntry
	lifetime time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(lifetime time.Duration) *Cache {
	newCache := Cache{
		Entry:    make(map[string]cacheEntry),
		lifetime: lifetime,
	}
	go newCache.readLoop()
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
	c.Entry[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.Entry[key]
	if !ok {
		return []byte{}, false
	}
	return entry.value, true
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.lifetime)
	for {
		currentTime := <-ticker.C
		for key, entry := range c.Entry {
			age := currentTime.Sub(entry.createdAt)
			if age >= c.lifetime {
				delete(c.Entry, key)
			}
		}
	}
}
