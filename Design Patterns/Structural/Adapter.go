package main

import "fmt"

// Target interface
// This is what our new system expects.
type Payment interface {
	Pay(amount int)
}

// Adaptee
// Old/third-party payment system.
type OldPaymentSystem struct{}

func (o *OldPaymentSystem) MakePayment(amount int) {
	fmt.Println("Payment of", amount, "made using old payment system")
}

// Adapter
// Converts OldPaymentSystem into Payment.
type PaymentAdapter struct {
	oldPayment *OldPaymentSystem
}

func (p *PaymentAdapter) Pay(amount int) {
	p.oldPayment.MakePayment(amount)
}

// Client
// Client simply means the code that wants to use the functionality.
func ProcessPayment(payment Payment, amount int) {
	payment.Pay(amount)
}

// func main() {

// 	oldSystem := &OldPaymentSystem{}

// 	adapter := &PaymentAdapter{
// 		oldPayment: oldSystem,
// 	}

// 	// Client only knows about Payment interface.
// 	ProcessPayment(adapter, 1000)
// }
