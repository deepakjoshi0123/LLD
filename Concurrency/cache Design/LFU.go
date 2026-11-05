package main

import (
	"container/list"
	"fmt"
	"sync"
)

type Entry struct {
	Key       string
	Value     int
	Frequency int
}

type LFUCache struct {
	capacity int

	// Key -> cache entry
	items map[string]*list.Element

	// Frequency -> list of entries with that frequency.
	// Within each list, most recently used is at the front.
	freqList map[int]*list.List

	// Minimum frequency currently present in the cache.
	minFreq int

	mu sync.Mutex
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		freqList: make(map[int]*list.List),
	}
}

// Get retrieves a value and increases its frequency.
func (c *LFUCache) Get(key string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, exists := c.items[key]
	if !exists {
		return 0, false
	}

	entry := element.Value.(*Entry)

	// Update the frequency and list position.
	c.increaseFrequency(element)

	return entry.Value, true
}

// Put inserts a new entry or updates an existing entry.
func (c *LFUCache) Put(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity <= 0 {
		return
	}

	// If the key already exists, update its value
	// and count this operation as an access.
	if element, exists := c.items[key]; exists {
		entry := element.Value.(*Entry)
		entry.Value = value
		c.increaseFrequency(element)
		return
	}

	// Evict an entry if the cache is full.
	if len(c.items) >= c.capacity {
		lruList := c.freqList[c.minFreq]
		lruElement := lruList.Back()
		lruEntry := lruElement.Value.(*Entry)

		delete(c.items, lruEntry.Key)
		lruList.Remove(lruElement)

		if lruList.Len() == 0 {
			delete(c.freqList, c.minFreq)
		}
	}

	// Insert the new entry with frequency 1.
	entry := &Entry{
		Key:       key,
		Value:     value,
		Frequency: 1,
	}

	if c.freqList[1] == nil { // new element always have frequency of 1 ----------
		c.freqList[1] = list.New()
	}

	element := c.freqList[1].PushFront(entry)
	c.items[key] = element

	// Every new entry starts at frequency 1.
	c.minFreq = 1
}

// Moves an entry from frequency f to frequency f+1.
func (c *LFUCache) increaseFrequency(element *list.Element) {
	entry := element.Value.(*Entry)
	oldFreq := entry.Frequency

	// Remove from the old frequency list.
	oldList := c.freqList[oldFreq]
	oldList.Remove(element)

	// If this was the last entry at the minimum frequency,
	// the minimum frequency increases.
	if oldList.Len() == 0 {
		delete(c.freqList, oldFreq)

		if c.minFreq == oldFreq {
			c.minFreq++
		}
	}

	// Increase the entry's frequency.
	entry.Frequency++

	// Add to the front of the new frequency list.
	if c.freqList[entry.Frequency] == nil {
		c.freqList[entry.Frequency] = list.New()
	}

	newElement := c.freqList[entry.Frequency].PushFront(entry)

	// Update the map because the list element has changed.
	c.items[entry.Key] = newElement
}

func main() {
	cache := NewLFUCache(3)

	cache.Put("A", 10)
	cache.Put("B", 20)
	cache.Put("C", 30)

	// Frequencies: A=1, B=1, C=1
	cache.Get("A") // A=2
	cache.Get("A") // A=3
	cache.Get("B") // B=2

	// Cache is full. C has the lowest frequency (1),
	// so C is evicted.
	cache.Put("D", 40)

	if value, ok := cache.Get("A"); ok {
		fmt.Println("A:", value)
	}

	if value, ok := cache.Get("C"); ok {
		fmt.Println("C:", value)
	} else {
		fmt.Println("C: evicted")
	}

	if value, ok := cache.Get("D"); ok {
		fmt.Println("D:", value)
	}
}
