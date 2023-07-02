package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// cacheObject is the cache object storing needed vals
type cacheObject struct {
	Value string `json:"val"`
	Count int    `json:"count"`
}

// cacheEngine is the cache engine
type cacheEngine struct {
	cap   int
	cache map[string]cacheObject
}

// NewCache is to initiate cache engine
func NewCache(cap int) *cacheEngine {
	return &cacheEngine{
		cap:   cap,
		cache: make(map[string]cacheObject),
	}
}

// Set is to assign val to key onto the cache
func (c *cacheEngine) Set(key string, val string) {
	log.Printf("set '%s' with '%s'...", key, val)

	mapVal, ok := c.cache[key]
	if ok {
		// when value found, then update count and value
		mapVal.Count++
		mapVal.Value = val
	} else {
		// when not found, do cap validation
		if len(c.cache) >= c.cap {
			log.Printf("cap reached...")

			// eliminate least accessed
			c.eliminateLeastAccessed()
		}
		mapVal = cacheObject{
			Value: val,
			Count: 1,
		}
	}

	c.cache[key] = mapVal
}

// eliminateLeastAccessed is to scan and
// remove cache with the least accessed count
func (c *cacheEngine) eliminateLeastAccessed() {
	leastKey, leastCount := "", -1
	for key := range c.cache {
		if leastCount < 0 || c.cache[key].Count < leastCount {
			leastCount = c.cache[key].Count
			leastKey = key
		}
	}

	log.Printf("eliminating '%s'...", leastKey)
	delete(c.cache, leastKey)
}

// Get is to get cache
func (c *cacheEngine) Get(key string) string {
	log.Printf("getting '%s'...", key)
	val, ok := c.cache[key]
	if ok {
		// when getting value, add access count
		val.Count++
		c.cache[key] = val
		return val.Value
	}
	return ""
}

// main
func main() {
	cache := NewCache(5)

	// for i := 0; i < 20; i++ {
	// 	keyI, valI := rand.Intn(10), rand.Intn(100)
	// 	cache.Set(fmt.Sprintf("key%d", keyI), fmt.Sprintf("val%d", valI))
	// }

	cache.Set("key1", "val")
	cache.Set("key2", "val")
	cache.Set("key1", "val")
	cache.Set("key3", "val")
	cache.Set("key2", "val")
	cache.Set("key4", "val")
	cache.Set("key3", "val")
	cache.Set("key5", "val")

	cache.Get("key5")
	cache.Get("key5")
	cache.Get("key5")
	cache.Get("key5")
	cache.Get("key1")
	cache.Get("key2")
	cache.Get("key3")
	cache.Get("key4")

	cache.Set("key4", "val")
	cache.Set("key6", "val")
	cache.Set("key5", "val")
	cache.Set("key7", "val")
	cache.Set("key6", "val")
	cache.Set("key8", "val")
	cache.Set("key7", "val")
	cache.Set("key9", "val")
	cache.Set("key8", "val")
	cache.Set("key10", "val")
	cache.Set("key9", "val")
	cache.Set("key10", "val")

	raw, _ := json.Marshal(cache.cache)
	// raw, _ := json.MarshalIndent(cache.cache, "", " ")
	fmt.Printf("%s", string(raw))
}
