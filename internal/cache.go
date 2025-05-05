package pokecache

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	cachemap  map[string]CacheEntry
	cachmutex sync.Mutex
	ttl       time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	rawdata   []byte
}

func NewCache(t time.Duration) *Cache {
	cache := &Cache{
		cachemap: make(map[string]CacheEntry),
		ttl:      t,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, data []byte) {
	c.cachmutex.Lock()
	defer c.cachmutex.Unlock()
	c.cachemap[key] = CacheEntry{
		rawdata:   data,
		createdAt: time.Now(),
	}
	log.Println("New Cache Entry created at ", c.cachemap[key].createdAt)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.cachemap[key]
	if !ok {
		log.Println("cant get data from cache its not here")
		return nil, false
	}
	log.Println("got data from cache")
	return val.rawdata, true
}

func (c *Cache) reapLoop() {
	for {
		time.Sleep(c.ttl)
		c.cachmutex.Lock()
		for key, val := range c.cachemap {
			if time.Since(val.createdAt) >= c.ttl {
				delete(c.cachemap, key)
				// log.Println("Some CacheEntry was deleted ")
			}
		}
		c.cachmutex.Unlock()
	}
}

// func (c *Cache) reapLoop() {
// 	ticker := time.NewTicker(c.ttl)
// 	defer ticker.Stop()

// 	for {
// 		<-ticker.C
// 		c.cachmutex.Lock()
// 		for key, val := range c.cachemap {
// 			if time.Since(val.createdAt) >= c.ttl {
// 				delete(c.cachemap, key)
// log.Println("CacheEntry deleted for key:", key)
// 			}
// 		}
// 		c.cachmutex.Unlock()
// 	}
// }
