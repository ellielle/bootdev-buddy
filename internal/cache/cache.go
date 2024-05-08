package cache

import (
	"sync"
	"time"
)

// cacheEntry is expecting JSON as a []byte
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache includes a mutex because maps aren't
// concurrency safe
type Cache struct {
	Entries map[string]cacheEntry
	mu      *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val *[]byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       *val,
	}
}

func (c *Cache) Get(entry string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, ok := c.Entries[entry]
	return data.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)

	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.Entries {
		if entry.createdAt.Before(now.Add(-interval)) {
			delete(c.Entries, key)
		}
	}

}
