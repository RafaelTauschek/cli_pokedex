package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	v  map[string]CacheEntry
	mu sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.v[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.v[key]
	c.mu.Unlock()

	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			<-ticker.C

			c.mu.Lock()

			for key, entry := range c.v {
				age := time.Since(entry.createdAt)

				if age > interval {
					delete(c.v, key)
				}
			}

			c.mu.Unlock()
		}
	}()
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		v: make(map[string]CacheEntry),
	}
	cache.reapLoop(interval)
	return cache
}
