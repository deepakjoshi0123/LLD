package main

import "fmt"

// Product interface
type Notificationn interface {
	Send(message string)
}

// Concrete Product 1
type SMSs struct{}

func (s *SMS) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

// Concrete Product 2
type Emaill struct{}

func (e *Email) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// Concrete Product 3
type WhatsAppp struct{}

func (w *WhatsApp) Send(message string) {
	fmt.Println("Sending WhatsApp:", message)
}

// Simple Factory
var factories = map[string]func() Notification{
	"sms": func() Notification {
		return &SMS{}
	},
	"email": func() Notification {
		return &Email{}
	},
	"whatsapp": func() Notification {
		return &WhatsApp{}
	},
}

func NotificationFactory_ocp(notificationType string) Notification {

	factory, ok := factories[notificationType]

	if !ok {
		return nil
	}

	return factory()
}

// func main() {

// 	notification := NotificationFactory("email")

// 	if notification != nil {
// 		notification.Send("Your order has been shipped")
// 	}
// }
