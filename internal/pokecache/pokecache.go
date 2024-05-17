package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mux     *sync.RWMutex
	entries map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mux:     &sync.RWMutex{},
		entries: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, value := range c.entries {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.entries, key)
		}
	}
}
