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

func NewCache(lifetime time.Duration) Cache {
	return Cache{
		Entry:    make(map[string]cacheEntry),
		lifetime: lifetime,
	}
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
