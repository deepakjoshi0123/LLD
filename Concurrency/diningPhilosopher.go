// The Problem Setup
// • The Scenario: Five philosophers sit around a circular table. In front of each philosopher is a bowl of spaghetti.
// • The Utensils: There are five chopsticks (or forks) placed between each adjacent pair of philosophers (one on the left and one on the right of each plate).
// • The Routine: Each philosopher continuously alternates between two states: thinking and eating.
// The Rules and Constraints
// • To eat, a philosopher must acquire both adjacent chopsticks (the one on their left and the one on their right).
// • A chopstick can only be picked up by one philosopher at a time.
// • If a neighboring philosopher is already using one of the chopsticks, the hungry philosopher must wait.
// After finishing a meal, the philosopher puts down both chopsticks so others can use them

// Core Synchronization Issues
// • Deadlock: If all five philosophers become hungry at the exact same time and pick up their left chopstick first, every chopstick is held, and each philosopher waits indefinitely for the right chopstick. No one can eat, and the system freezes in a circular wait.
// • Starvation: A situation where a single philosopher is continuously blocked from acquiring both chopsticks because neighboring processes repeatedly beat them to the shared resources.

package main

import (
	"fmt"
	"sync"
	"time"
)

// ????????????????????????????????????????????????????????????????????????????
const philosophers = 5

type Philosopher struct {
	id    int
	left  *sync.Mutex
	right *sync.Mutex
}

func (p Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {

		// THINKING
		fmt.Printf("Philosopher %d is thinking\n", p.id)
		time.Sleep(500 * time.Millisecond)

		// HUNGRY
		fmt.Printf("Philosopher %d is hungry\n", p.id)

		// Always acquire the lower-numbered chopstick first.
		// This prevents circular waiting.
		first := p.left
		second := p.right

		if p.id%2 == 0 {
			first = p.right
			second = p.left
		}

		// PICK UP FIRST CHOPSTICK
		first.Lock()
		fmt.Printf("Philosopher %d picked up first chopstick\n", p.id)

		// PICK UP SECOND CHOPSTICK
		second.Lock()
		fmt.Printf("Philosopher %d picked up second chopstick\n", p.id)

		// EATING
		fmt.Printf("🍝 Philosopher %d is eating\n", p.id)
		time.Sleep(1 * time.Second)

		// PUT DOWN CHOPSTICKS
		second.Unlock()
		first.Unlock()

		fmt.Printf("Philosopher %d put down both chopsticks\n", p.id)
	}
}

func main() {

	// Five chopsticks
	var chopsticks [philosophers]sync.Mutex

	var wg sync.WaitGroup

	for i := 0; i < philosophers; i++ {

		left := &chopsticks[i]
		right := &chopsticks[(i+1)%philosophers]

		p := Philosopher{
			id:    i,
			left:  left,
			right: right,
		}

		wg.Add(1)
		go p.dine(&wg)
	}

	wg.Wait()

	fmt.Println("Dinner finished")
}
