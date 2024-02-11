package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mux        sync.Mutex
	duration   time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cacheEntry[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cacheEntry[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(ticker *time.Ticker) {
	for {
		<-ticker.C
		c.mux.Lock()
		for key, val := range c.cacheEntry {
			if time.Since(val.createdAt) > c.duration {
				delete(c.cacheEntry, key)
			}
		}
		c.mux.Unlock()
	}
}

func NewCache(duration time.Duration) Cache {
	newCache := Cache{
		cacheEntry: make(map[string]cacheEntry),
		mux:        sync.Mutex{},
		duration:   duration,
	}
	ticker := time.NewTicker(duration)
	go newCache.reapLoop(ticker)

	return newCache
}
