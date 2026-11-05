// token bucket
// -------------->> check Prateek's implmentation

// Suppose we configure:
// - Bucket capacity: 5 tokens
// - Refill rate: 2 tokens per second
// - Each request consumes: 1 token

// =======================================================================================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	mu         sync.Mutex
	capacity   float64
	tokens     float64
	refillRate float64
	lastRefill time.Time
}

func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()

	b.tokens += elapsed * b.refillRate
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}
	b.lastRefill = now

	if b.tokens >= 1 {
		b.tokens--
		return true
	}

	return false
}

type RateLimiter struct {
	mu      sync.Mutex
	buckets map[string]*TokenBucket
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		buckets: make(map[string]*TokenBucket),
	}
}

func (r *RateLimiter) AddUser(
	userID string,
	capacity, refillRate float64,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.buckets[userID] = NewTokenBucket(capacity, refillRate)
}

func (r *RateLimiter) Allow(userID string) bool {
	r.mu.Lock()
	bucket, exists := r.buckets[userID]
	r.mu.Unlock()

	if !exists {
		return false // User has not been configured.
	}

	return bucket.Allow()
}

func main() {
	limiter := NewRateLimiter()

	// Free user: 5 tokens, refills 1 token/sec.
	limiter.AddUser("alice", 5, 1)

	// Premium user: 10 tokens, refills 3 tokens/sec.
	limiter.AddUser("bob", 10, 3)

	for i := 1; i <= 6; i++ {
		fmt.Printf("Alice request %d: %v\n",
			i, limiter.Allow("alice"))
	}

	for i := 1; i <= 11; i++ {
		fmt.Printf("Bob request %d: %v\n",
			i, limiter.Allow("bob"))
	}
}
