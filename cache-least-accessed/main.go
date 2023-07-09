package main

import (
	"log"
)

// cacheEngine is the cache engine
type cacheEngine struct {
}

// NewCache is to initiate cache engine
func NewCache(cap int) *cacheEngine {
	return &cacheEngine{}
}

// Set is to assign val to key onto the cache
func (c *cacheEngine) Set(key string, val string) {
	log.Printf("set '%s' with '%s'...", key, val)

}

// Get is to get cache
func (c *cacheEngine) Get(key string) string {
	log.Printf("getting '%s'...", key)
	// implement
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

	// raw, _ := json.Marshal(cache.cache)
	// raw, _ := json.MarshalIndent(cache.cache, "", " ")
	// fmt.Printf("%s", string(raw))
}
