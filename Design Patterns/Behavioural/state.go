// Why does ATM have the same function names?

// Because we want the user of the ATM to simply do:

// atm.insert()
// atm.withdrawcash()
// atm.removecard()
// atm.presscancel()

package main

import "fmt"

// State
type ATMState interface {
	insert()
	withdrawcash()
	removecard()
	presscancel()
}

// Context
type ATM struct {
	state ATMState
}

func (a *ATM) setState(state ATMState) {
	a.state = state
}

func (a *ATM) insert() {
	a.state.insert()
}

func (a *ATM) withdrawcash() {
	a.state.withdrawcash()
}

func (a *ATM) removecard() {
	a.state.removecard()
}

func (a *ATM) presscancel() {
	a.state.presscancel()
}

// --------------------
// Card Not Inserted
// --------------------

type CardNotInsertedState struct {
	atm *ATM
}

func (c *CardNotInsertedState) insert() {
	fmt.Println("Card inserted")
	c.atm.setState(&CardInsertedState{atm: c.atm})
}

func (c *CardNotInsertedState) withdrawcash() {
	fmt.Println("Please insert your card first")
}

func (c *CardNotInsertedState) removecard() {
	fmt.Println("No card to remove")
}

func (c *CardNotInsertedState) presscancel() {
	fmt.Println("Nothing to cancel")
}

// --------------------
// Card Inserted
// --------------------

type CardInsertedState struct {
	atm *ATM
}

func (c *CardInsertedState) insert() {
	fmt.Println("Card already inserted")
}

func (c *CardInsertedState) withdrawcash() {
	fmt.Println("Cash withdrawn")
}

func (c *CardInsertedState) removecard() {
	fmt.Println("Card removed")
	c.atm.setState(&CardNotInsertedState{atm: c.atm})
}

func (c *CardInsertedState) presscancel() {
	fmt.Println("Transaction cancelled")
	c.atm.setState(&CardNotInsertedState{atm: c.atm})
}

// --------------------
// Main
// --------------------

// func main() {

// 	atm := &ATM{}

// 	// Initial state
// 	atm.setState(&CardNotInsertedState{atm: atm})

// 	atm.withdrawcash()

// 	atm.insert()

// 	atm.withdrawcash()

// 	atm.presscancel()

// 	atm.withdrawcash()
// }
