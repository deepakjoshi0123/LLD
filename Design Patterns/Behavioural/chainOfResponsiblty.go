package main

import "fmt"

type Handler interface {
	SetNext(Handler)
	Handle(int)
}

// Base handler avoids repeating next-handler logic.
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(next Handler) {
	b.next = next
}

func (b *BaseHandler) Handle(amount int) {
	if b.next != nil {
		b.next.Handle(amount)
	}
}

type Manager struct {
	BaseHandler
}

func (m *Manager) Handle(amount int) {
	if amount <= 1000 {
		fmt.Println("Manager handled the request")
		return
	}

	m.BaseHandler.Handle(amount)
}

type Director struct {
	BaseHandler
}

func (d *Director) Handle(amount int) {
	if amount <= 5000 {
		fmt.Println("Director handled the request")
		return
	}

	d.BaseHandler.Handle(amount)
}

type CEO struct {
	BaseHandler
}

func (c *CEO) Handle(amount int) {
	if amount <= 10000 {
		fmt.Println("CEO handled the request")
		return
	}

	fmt.Println("Request cannot be handled")
}

// func main() {
// 	manager := &Manager{}
// 	director := &Director{}
// 	ceo := &CEO{}

// 	// Build the chain:
// 	// Manager → Director → CEO
// 	manager.SetNext(director)
// 	director.SetNext(ceo)

// 	// Send requests through the first handler.
// 	manager.Handle(800)
// 	manager.Handle(3000)
// 	manager.Handle(8000)
// 	manager.Handle(15000)
// }
