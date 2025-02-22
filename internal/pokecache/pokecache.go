package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
    cache   map[string]cacheEntry
    mu      sync.Mutex
    done    chan bool
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
    defer c.mu.Unlock()
    if val, ok := c.cache[key]; ok {
        return val.val, true
    }

    return nil, false
}

func (c *Cache) Close() {
    c.done <- true
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    defer ticker.Stop()  // Clean up ticker when we exit

    for {
        select {
        case <-ticker.C:
            c.mu.Lock()
            for k, v := range c.cache {
                if time.Since(v.createdAt) > interval {
                    delete(c.cache, k)
                }
            }
            c.mu.Unlock()
        case <-c.done:
            return
        }
    }
}

func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        cache: make(map[string]cacheEntry),
        done:  make(chan bool),
    }

    go c.reapLoop(interval)
    return c
}
