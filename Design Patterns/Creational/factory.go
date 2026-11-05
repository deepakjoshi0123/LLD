package main

import "fmt"

// it's for only one type of product
// and anyway we will break ocp in notifcation factory as soon as new implementation comes in

// Product interface
type Notification interface {
	Sendd(message string)
}

// Concrete Product 1
type SMS struct{}

func (s *SMS) Sendd(message string) {
	fmt.Println("Sending SMS:", message)
}

// Concrete Product 2
type Email struct{}

func (e *Email) Sendd(message string) {
	fmt.Println("Sending Email:", message)
}

// Concrete Product 3
type WhatsApp struct{}

func (w *WhatsApp) Sendd(message string) {
	fmt.Println("Sending WhatsApp:", message)
}

// Factory
func NotificationFactory(notificationType string) Notification {

	switch notificationType {
	case "sms":
		return &SMS{}

	case "email":
		return &Email{}

	case "whatsapp":
		return &WhatsApp{}

	default:
		return nil
	}
}

// func main() {

// 	notification := NotificationFactory("email")

// 	notification.Send("Your order has been shipped")
// }
