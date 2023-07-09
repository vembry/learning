package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// node is the struct to store cache object
type node struct {
	Value string `json:"value"`
}

// cacheEngine is the cache engine
type cacheEngine struct {
	cap   int
	cache map[string]*node
}

// NewCache is to initiate cache engine
func NewCache(cap int) *cacheEngine {
	return &cacheEngine{
		cap:   cap,
		cache: make(map[string]*node),
	}
}

// Set is to assign val to key onto the cache
func (c *cacheEngine) Set(key string, val string) {
	log.Printf("set '%s' with '%s'...", key, val)

	// assign key-val to the cache
	c.cache[key] = &node{
		Value: val,
	}
}

// Get is to get cache
func (c *cacheEngine) Get(key string) string {
	log.Printf("getting '%s'...", key)

	// get from cache
	if val, ok := c.cache[key]; ok {
		return val.Value
	}

	// return empty when not found
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
