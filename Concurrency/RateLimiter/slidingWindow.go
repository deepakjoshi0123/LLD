// sliding window
// <<<<<<<<<<<<<<<<<<<<<<================>>>>>>>>>>>>>>>>>>>>>

package main

import (
	"sync"
	"time"
)

type Transaction struct {
	Amount int64
	Time   time.Time
}

type User struct {
	DailyLimit   int64
	Transactions []Transaction
	mu           sync.Mutex
}

type PaymentLimiter struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewPaymentLimiter() *PaymentLimiter {
	return &PaymentLimiter{
		users: make(map[string]*User),
	}
}

func (p *PaymentLimiter) AddUser(userID string, limit int64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.users[userID] = &User{
		DailyLimit: limit,
	}
}

func (p *PaymentLimiter) AllowPayment(
	userID string,
	amount int64,
) bool {
	if amount <= 0 {
		return false
	}

	p.mu.RLock()
	user, exists := p.users[userID]
	p.mu.RUnlock()

	if !exists {
		return false
	}

	user.mu.Lock()
	defer user.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-24 * time.Hour)

	// Remove transactions older than the rolling 24-hour window.
	valid := user.Transactions[:0]
	var total int64

	for _, tx := range user.Transactions {
		if tx.Time.After(windowStart) {
			valid = append(valid, tx)
			total += tx.Amount
		}
	}
	user.Transactions = valid

	// Prevent the transaction from exceeding the user's limit.
	if amount > user.DailyLimit-total {
		return false
	}

	// Record the accepted transaction.
	user.Transactions = append(user.Transactions, Transaction{
		Amount: amount,
		Time:   now,
	})

	return true
}

// func main() {
// 	limiter := NewPaymentLimiter()

// 	// Different limits for different users, in rupees.
// 	limiter.AddUser("userA", 10_000)
// 	limiter.AddUser("userB", 50_000)
// 	limiter.AddUser("userC", 100_000)

// 	fmt.Println(limiter.AllowPayment("userA", 3_000))  // true
// 	fmt.Println(limiter.AllowPayment("userA", 8_000))  // false
// 	fmt.Println(limiter.AllowPayment("userA", 7_000))  // true
// 	fmt.Println(limiter.AllowPayment("userB", 40_000)) // true
// 	fmt.Println(limiter.AllowPayment("userC", 90_000)) // true
// }
