package main

import "fmt"

// Strategy
type PaymentStrategy interface {
	Pay(amount float64)
}

// Concrete Strategy 1
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) {
	fmt.Println("Paid", amount, "using Credit Card")
}

// Concrete Strategy 2
type UPI struct{}

func (u *UPI) Pay(amount float64) {
	fmt.Println("Paid", amount, "using UPI")
}

// Concrete Strategy 3
type PayPal struct{}

func (p *PayPal) Pay(amount float64) {
	fmt.Println("Paid", amount, "using PayPal")
}

// Context
type Payment struct {
	strategy PaymentStrategy
}

func (p *Payment) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *Payment) MakePayment(amount float64) {
	p.strategy.Pay(amount)
}

// func main() {

// 	payment := &Payment{}

// 	// Use Credit Card
// 	payment.SetStrategy(&CreditCard{})
// 	payment.MakePayment(1000)

// 	// Change strategy to UPI
// 	payment.SetStrategy(&UPI{})
// 	payment.MakePayment(2000)

// 	// Change strategy to PayPal
// 	payment.SetStrategy(&PayPal{})
// 	payment.MakePayment(3000)
// }
