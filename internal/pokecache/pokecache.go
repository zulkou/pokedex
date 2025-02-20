package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
    cache   map[string]cacheEntry
    mu      sync.Mutex
}

type cacheEntry struct {
    createdAt   time.Time
    val         []byte
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    c.cache[key] = cacheEntry{
        createdAt: time.Now(),
        val: val,
    }
    c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    if val, ok := c.cache[key]; ok {
        return val.val, true
    }

    defer c.mu.Unlock()
    return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)

    for {
        <-ticker.C

        c.mu.Lock()
        for k, v := range c.cache {
            if (time.Now().Sub(v.createdAt) > interval) {
                delete(c.cache, k)
            }
        }
        c.mu.Unlock()
    }
}

func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        cache: make(map[string]cacheEntry),
    }

    go c.reapLoop(interval)
    return c
}
