package main

import (
	"fmt"
	"log"
)

type CacheProvider interface {
	// Set is to set cache key-value pair
	Set(key string, val string)

	// Get is to get cache value by key
	Get(key string) string
}

type Node struct {
	Key  string
	Val  string
	Next *Node
	Prev *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("key:%s, val:%s", n.Key, n.Val)
}

type LRUCache struct {
	cacheMap map[string]*Node
	cap      int
	head     *Node
	tail     *Node
}

// NewLRUCache is to initialize LRUCache
func NewLRUCache(cap int) *LRUCache {

	head := &Node{
		Key: "head",
		Val: "head",
	}
	tail := &Node{
		Key: "tail",
		Val: "tail",
	}

	head.Prev = tail
	tail.Next = head

	return &LRUCache{
		cacheMap: map[string]*Node{},
		cap:      cap,
		head:     head,
		tail:     tail,
	}
}

// Set is to set cache key-value pair for LRUCache
func (c *LRUCache) Set(key string, val string) {
	log.Printf("inserting key=%s, val=%s...", key, val)

	// construct cache node
	newNode := &Node{
		Key: key,
		Val: val,
	}

	if val, ok := c.cacheMap[key]; ok {
		// when cache exists, get it and replace newNode
		newNode = c.takeOutNode(val)
		delete(c.cacheMap, key)
	} else if len(c.cacheMap) == c.cap {
		// when len reach the cap, then remove least-used
		val := c.takeOutNode(c.tail.Next)
		delete(c.cacheMap, val.Key)
	}

	c.cacheMap[key] = newNode
	// set new node on top
	c.setOnTop(newNode)

	print(*c)
}

// setOnTop is to place node on top prior to head
func (c *LRUCache) setOnTop(node *Node) {
	// make newNode to replace Head
	node.Prev = c.head.Prev

	// fix the chain between head and head's prev
	c.head.Prev.Next = node

	// shift head to be in-front of newNode
	node.Next = c.head
	c.head.Prev = node
}

// takeOutNode is to gracefully takeout node
func (c *LRUCache) takeOutNode(node *Node) *Node {

	// fix chain that covered by selected node
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	node.Next = nil
	node.Prev = nil

	return node
}

// Get is to get cache value by key for LRUCache
func (c *LRUCache) Get(key string) string {
	log.Printf("getting key=%s...", key)
	if val, ok := c.cacheMap[key]; ok {
		val = c.takeOutNode(val)
		c.setOnTop(val)

		print(*c)
		return val.Val
	}
	return ""
}

func main() {
	lruCache := NewLRUCache(6)
	cache := CacheProvider(lruCache)

	cache.Set("key1", "asd")
	cache.Set("key2", "asd")
	cache.Set("key1", "asd")
	cache.Set("key3", "asd")
	cache.Set("key4", "asd")
	cache.Set("key5", "asd")
	cache.Get("key2")
	cache.Set("key6", "asd")
	cache.Set("key7", "asd")
	cache.Get("key2")
	cache.Set("key8", "asd")
	cache.Set("key9", "asd")
	cache.Set("key10", "asd")
	// print(*lruCache)
	cache.Set("key1", "asd")
	// print(*lruCache)

}

// print is to debug with less logs
func print(cache LRUCache) {
	fmt.Println("----------------------------------------")
	for x := cache.tail.Next; x.Next != nil; x = x.Next {
		// print curr
		curr := fmt.Sprintf("%s", x)
		fmt.Println(curr)
	}
	fmt.Println("----------------------------------------")
}

// printDebug is to debut cache :p
func printDebug(cache LRUCache) {
	fmt.Println("----------------------------------------")
	for x := cache.tail.Next; x.Next != nil; x = x.Next {

		// print curr
		curr := fmt.Sprintf("- %s", x)
		fmt.Println(curr)

		// print prev
		prev := "-- prev=nil"
		if x.Prev != nil {
			prev = fmt.Sprintf("-- prev=%s", x.Prev)
		}
		fmt.Println(prev)

		// print next
		next := "-- next=nil"
		if x.Next != nil {
			next = fmt.Sprintf("-- next=%s", x.Next)
		}
		fmt.Println(next)
		fmt.Println()
	}
	fmt.Println("----------------------------------------")
}
