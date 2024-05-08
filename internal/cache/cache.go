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

// NewCache creates a new Cache to be used
// during the session
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

// Add creates a new entry in the cache
func (c *Cache) Add(key string, val *[]byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       *val,
	}
}

// Get attempts to retrieve an entry from the cache
func (c *Cache) Get(entry string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, ok := c.Entries[entry]
	return data.val, ok
}

// reapLoop controls the removal of old cache entries
// by reaping entries older than time.Now() + interval
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)

	}
}

// reap deletes old entries from the cache
func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.Entries {
		if entry.createdAt.Before(now.Add(-interval)) {
			delete(c.Entries, key)
		}
	}

}
