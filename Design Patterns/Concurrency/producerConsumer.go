// The Problem
// The core challenge is to ensure that the producer and consumer do not access the shared buffer at the same time, while respecting the physical limits of the buffer. This requires solving three specific sub-problems:
// 1. The Full Buffer Constraint (Overflow): If the buffer is completely full, the producer must be forced to wait. It cannot add new data until the consumer removes at least one item to free up space.
// 2. The Empty Buffer Constraint (Underflow): If the buffer is completely empty, the consumer must be forced to wait. It cannot retrieve data until the producer inserts at least one new item.
// 3. The Critical Section Problem (Race Conditions): The buffer is a shared resource. If a producer is trying to insert an item at the exact same moment a consumer is trying to remove one, they may overwrite the same memory address or corrupt the data tracking variables (like the buffer count). Only one process must be allowed to modify the buffer at any given millisecond.

// What is sync.Cond?
// sync.Cond means Condition Variable.
// It basically allows a goroutine to say:
// "I can't continue right now. Wake me up when some condition becomes true."

package main

import (
	"fmt"
	"sync"
	"time"
)

type Buffer struct {
	items    []int
	capacity int

	mu       sync.Mutex
	notFull  *sync.Cond
	notEmpty *sync.Cond
}

// Create a new buffer
func NewBuffer(capacity int) *Buffer {
	b := &Buffer{
		items:    make([]int, 0, capacity),
		capacity: capacity,
	}

	b.notFull = sync.NewCond(&b.mu)  // Wait() already knows which locker it is associated with
	b.notEmpty = sync.NewCond(&b.mu)

	return b
}

// Producer puts an item into the buffer
func (b *Buffer) Produce(item int) {

	b.mu.Lock()
	defer b.mu.Unlock()

	// If buffer is full, producer waits
	for len(b.items) == b.capacity {
		fmt.Println("Buffer is full. Producer is waiting...")
		b.notFull.Wait()
	}

	// Add item
	b.items = append(b.items, item)

	fmt.Println("Produced:", item)

	// Tell consumer that buffer is not empty
	b.notEmpty.Signal()
}

// This is the most important thing to understand about sync.Cond.Wait().
// When you call:
// b.notFull.Wait()

// Go atomically does two things:
// 1. Releases b.mu
// 2. Puts the goroutine to sleep
// Later, when it is signaled, it:
// 3. Wakes the goroutine
// 4. Re-acquires b.mu
// 5. Returns from Wait()

// Consumer takes an item from the buffer
func (b *Buffer) Consume() int {

	b.mu.Lock()
	defer b.mu.Unlock()

	// If buffer is empty, consumer waits
	for len(b.items) == 0 {
		fmt.Println("Buffer is empty. Consumer is waiting...")
		b.notEmpty.Wait()
	}

	// Take first item
	item := b.items[0]

	// Remove first item
	b.items = b.items[1:]

	fmt.Println("Consumed:", item)

	// Tell producer that buffer is not full
	// 	If the consumer goroutine hasn't reached Wait() yet, those Signal() calls effectively do nothing.
	// They are not stored for later.

	b.notFull.Signal()

	return item
}

// Producer goroutine
func producer(buffer *Buffer) {

	for i := 1; i <= 10; i++ {

		buffer.Produce(i)

		time.Sleep(500 * time.Millisecond)
	}
}

// Consumer goroutine
func consumer(buffer *Buffer) {

	for i := 1; i <= 10; i++ {

		buffer.Consume()

		time.Sleep(1 * time.Second)
	}
}

// func main() {

// 	// Buffer can hold only 3 items
// 	buffer := NewBuffer(3)

// 	// Start producer
// 	go producer(buffer)

// 	// Start consumer
// 	go consumer(buffer)

// 	// Wait so goroutines can finish
// 	// 	time.Sleep(15 * time.Second) is only there to keep main() alive. It is not part of the Producer–Consumer logic.
// 	// Remember: when main() finishes, the entire Go program exits, even if goroutines are still running.
// 	time.Sleep(15 * time.Second)

// 	fmt.Println("Program finished")
// }

// 1- producer consumer problem / multiple producer,consumer(extended question)
// 	- notify_all threads
