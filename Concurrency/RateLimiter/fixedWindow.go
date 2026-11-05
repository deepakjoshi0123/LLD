// fixed window
// ==============
// A Fixed Window Rate Limiter allows a maximum number of requests during a fixed time interval.

// For example, allow 5 requests per 10 seconds.
// - At the beginning of each 10-second window, the counter resets to zero.
// - Each incoming request increments the counter.
// - If the counter reaches 5, additional requests are rejected until the next window starts.

// =======================================================================================================================================================

package main

import (
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	mu          sync.Mutex
	limit       int
	window      time.Duration
	count       int
	windowStart time.Time
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:       limit,
		window:      window,
		windowStart: time.Now(),
	}
}

func (l *FixedWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	if now.Sub(l.windowStart) >= l.window {
		// now()- windowsstart >= 10 seconds(iwndown length)
		l.count = 0
		l.windowStart = now
	}

	if l.count >= l.limit {
		return false
	}

	l.count++
	return true
}

// func main() {
// 	// Each user gets an independent limiter.
// 	users := map[string]*FixedWindowLimiter{
// 		"Alice":   NewFixedWindowLimiter(3, 10*time.Second),
// 		"Bob":     NewFixedWindowLimiter(2, 10*time.Second),
// 		"Charlie": NewFixedWindowLimiter(4, 10*time.Second),
// 	}

// 	// Simulate requests from different users.
// 	requests := []string{
// 		"Alice", "Alice", "Bob", "Alice", "Bob",
// 		"Charlie", "Alice", "Bob", "Charlie", "Charlie",
// 	}

// 	for i, user := range requests {
// 		allowed := users[user].Allow()

// 		fmt.Printf(
// 			"Request %2d | User: %-7s | Allowed: %v\n",
// 			i+1, user, allowed,
// 		)
// 	}
// }
