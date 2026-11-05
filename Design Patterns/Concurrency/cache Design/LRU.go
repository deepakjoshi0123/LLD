// LRU
// -------
package main

import (
	"container/list"
	"fmt"
)

type Entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	items    map[int]*list.Element // understand how this Element helping in type assertion .........
	order    *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity < 0 {
		capacity = 0
	}

	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		order:    list.New(),
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	element, exists := c.items[key]
	if !exists {
		return 0, false
	}

	// Mark the accessed entry as most recently used.
	c.order.MoveToFront(element)

	entry := element.Value.(*Entry)
	return entry.value, true
}

func (c *LRUCache) Put(key, value int) {
	if c.capacity == 0 {
		return
	}

	// Update an existing entry.
	if element, exists := c.items[key]; exists {
		element.Value.(*Entry).value = value
		c.order.MoveToFront(element)
		return
	}

	// Evict the least recently used entry if full.
	if c.order.Len() >= c.capacity {
		lru := c.order.Back()
		entry := lru.Value.(*Entry)

		delete(c.items, entry.key) // delete from map
		c.order.Remove(lru)        // delete from DLL
	}

	// Insert new entry at the front.
	entry := &Entry{
		key:   key,
		value: value,
	}

	element := c.order.PushFront(entry)
	c.items[key] = element
}

func main() {
	cache := NewLRUCache(3)

	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)

	value, ok := cache.Get(1)
	fmt.Println(value, ok) // 10 true

	cache.Put(4, 40) // Evicts key 2

	_, ok = cache.Get(2)
	fmt.Println(ok) // false

	value, ok = cache.Get(3)
	fmt.Println(value, ok) // 30 true

	cache.Put(5, 50) // Evicts key 1

	_, ok = cache.Get(1)
	fmt.Println(ok) // false
}
