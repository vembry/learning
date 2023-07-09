package main

import (
	"fmt"
	"log"
	"math"
)

// node is the struct to store cache object
type node struct {
	Key   string
	Value string
	Count int
	Next  *node
	Prev  *node
}

// cacheEngine is the cache engine
type cacheEngine struct {
	cap   int
	cache map[string]*node
	head  *node
	tail  *node
}

// NewCache is to initiate cache engine
func NewCache(cap int) *cacheEngine {

	// init head and tail
	head := &node{Key: "head", Value: "head", Count: math.MaxInt}
	tail := &node{Key: "tail", Value: "tail", Count: math.MinInt}
	head.Prev = tail
	tail.Next = head

	return &cacheEngine{
		cap:   cap,
		cache: make(map[string]*node),
		head:  head,
		tail:  tail,
	}
}

// Set is to assign val to key onto the cache
func (c *cacheEngine) Set(key string, val string) {
	log.Printf("set key='%s' with val='%s'...", key, val)

	newNode := &node{
		Key:   key,
		Value: val,
		Count: 1,
	}

	if val, ok := c.cache[key]; ok {
		// validate if key exists
		val.Value = newNode.Value
		val.Count++

		shiftNodeUp(val)

		newNode = val
	} else {
		if len(c.cache) == c.cap {
			// validate capacity threshold
			log.Printf("reached cap %d", c.cap)
			temp := takeOutNode(c.tail.Next)
			delete(c.cache, temp.Key)
		}
		// re-chain the nodes with new node
		newNode.Next = c.tail.Next
		newNode.Prev = c.tail
		c.tail.Next = newNode
	}

	// assign key-val to the cache
	c.cache[key] = newNode

	print(c)
}

// shiftNodeUp is to shift up nodes
func shiftNodeUp(curr *node) {
	for curr.Next != nil {
		if curr.Count > curr.Next.Count {
			next := curr.Next

			// fix outer nodes chain of
			// before-curr node and after-next node
			if next.Next != nil {
				next.Next.Prev = curr
			}
			if curr.Prev != nil {
				curr.Prev.Next = next
			}

			// fix curr and next chain
			curr.Next = next.Next
			next.Next = curr

			next.Prev = curr.Prev
			curr.Prev = next
		} else {
			break
		}
	}
}

// takeOutNode is to extract out node out linked list
func takeOutNode(curr *node) *node {
	curr.Next.Prev = curr.Prev
	curr.Prev.Next = curr.Next

	curr.Next = nil
	curr.Prev = nil
	return curr
}

// Get is to get cache
func (c *cacheEngine) Get(key string) string {
	log.Printf("getting '%s'...", key)

	// get from cache
	if val, ok := c.cache[key]; ok {
		val.Count++
		shiftNodeUp(val)
		return val.Value
	}

	// return empty when not found
	return ""
}

// main
func main() {
	cache := NewCache(5)
	print(cache)

	// for i := 0; i < 20; i++ {
	// 	keyI, valI := rand.Intn(10), rand.Intn(100)
	// 	cache.Set(fmt.Sprintf("key%d", keyI), fmt.Sprintf("val%d", valI))
	// }

	cache.Set("key1", "val")
	cache.Set("key2", "val")
	cache.Set("key1", "val")
	cache.Set("key3", "val")
	cache.Set("key3", "val1")
	cache.Set("key3", "val2")
	cache.Set("key3", "val3")
	cache.Set("key3", "val4")
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
	print(cache)

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

	print(cache)
}

func print(cache *cacheEngine) {
	fmt.Println("==============================================================================")
	fmt.Println("printing")
	fmt.Println("==============================================================================")
	for x := cache.tail; x != nil; x = x.Next {
		fmt.Println(
			fmt.Sprintf("key=%s, val=%s, count=%d", x.Key, x.Value, x.Count),
		)
	}
	fmt.Println("==============================================================================")
	fmt.Println("closed")
	fmt.Println("==============================================================================")
}
