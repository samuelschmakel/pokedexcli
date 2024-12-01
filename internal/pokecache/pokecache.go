package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data map[string]cacheEntry
	mu   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

var myCache Cache

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Data: make(map[string]cacheEntry),
		mu:   sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if the map is initialized
	if c.Data == nil {
		c.Data = make(map[string]cacheEntry)
	}

	entry := cacheEntry{}
	entry.val = value
	entry.createdAt = time.Now()
	c.Data[key] = entry
	return
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if value, ok := c.Data[key]; ok {
		return value.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for now := range ticker.C {
		c.mu.Lock()
		for key, entry := range c.Data {
			if now.Sub(entry.createdAt) > interval {
				delete(c.Data, key)
			}
		}
		c.mu.Unlock()
	}
}
