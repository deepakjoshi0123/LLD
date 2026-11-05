// leaky bucket
// -------------

// The Leaky Bucket algorithm processes requests at a controlled, steady rate. Requests enter a bucket (usually a queue), and they are processed at a fixed rate. If the bucket is full, new requests are rejected.

// Unlike Token Bucket, which allows bursts using accumulated tokens, Leaky Bucket smooths incoming traffic.

// ``````````````````````````````````````
// Buffer contains A, B, C.
// Request D arrives → rejected.
// The ticker fires and the worker removes A.
// Request E arrives → accepted because space is available.
// ``````````````````````````````````````
// ======================================================================================================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	queue chan string
}

func NewLeakyBucket(capacity int, interval time.Duration) *LeakyBucket {
	b := &LeakyBucket{
		queue: make(chan string, capacity),
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			select {
			case request := <-b.queue:
				fmt.Println("Processed:", request)
			default:
			}
		}
	}()

	return b
}

func (b *LeakyBucket) Allow(request string) bool {
	select {
	case b.queue <- request:
		return true
	default:
		return false // Queue is full.
	}
}

type RateLimiter struct {
	mu    sync.Mutex
	users map[string]*LeakyBucket
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		users: make(map[string]*LeakyBucket),
	}
}

func (r *RateLimiter) AddUser(
	userID string,
	capacity int,
	interval time.Duration,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[userID]; exists {
		return
	}

	r.users[userID] = NewLeakyBucket(capacity, interval)
}

func (r *RateLimiter) Allow(userID, request string) bool {
	r.mu.Lock()
	bucket, exists := r.users[userID]
	r.mu.Unlock()

	if !exists {
		return false // Unknown user.
	}

	return bucket.Allow(request) //specified bucket for each user
}

// func main() {
// 	limiter := NewRateLimiter()

// 	// Alice: queue capacity 3, process one request/sec.
// 	limiter.AddUser("alice", 3, time.Second)

// 	// Bob: queue capacity 5, process one request every 500ms.
// 	limiter.AddUser("bob", 5, 500*time.Millisecond)

// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Alice:", i,
// 			limiter.Allow("alice", fmt.Sprintf("A-%d", i)))
// 	}

// 	for i := 1; i <= 7; i++ {
// 		fmt.Println("Bob:", i,
// 			limiter.Allow("bob", fmt.Sprintf("B-%d", i)))
// 	}

// 	time.Sleep(4 * time.Second)
// }
