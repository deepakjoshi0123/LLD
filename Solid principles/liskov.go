package main

import "fmt"

// =====================================================
// WITHOUT LSP
// =====================================================

// Bird has a Fly() method
type BirdWithoutLSP interface {
	Fly()
}

// Sparrow can fly
type SparrowWithoutLSP struct{}

func (s SparrowWithoutLSP) Fly() {
	fmt.Println("Sparrow is flying")
}

// Penguin cannot fly
type PenguinWithoutLSP struct{}

func (p PenguinWithoutLSP) Fly() {
	// Problem!
	fmt.Println("Penguin cannot fly")
}

func makeBirdFly(b BirdWithoutLSP) {
	b.Fly()
}

// =====================================================
// WITH LSP
// =====================================================

// Bird only contains behavior common to all birds
type Bird struct {
	Name string
}

func (b Bird) Eat() {
	fmt.Println(b.Name, "is eating")
}

// FlyingBird represents birds that can actually fly
type FlyingBird interface {
	Fly()
}

// Sparrow can fly
type Sparrow struct {
	Bird
}

func (s Sparrow) Fly() {
	fmt.Println(s.Name, "is flying")
}

// Penguin is a bird, but it does not implement Fly()
type Penguin struct {
	Bird
}

func main() {

	fmt.Println("========== WITHOUT LSP ==========")

	sparrow1 := SparrowWithoutLSP{}
	penguin1 := PenguinWithoutLSP{}

	makeBirdFly(sparrow1)
	makeBirdFly(penguin1)

	fmt.Println()

	fmt.Println("========== WITH LSP ==========")

	sparrow2 := Sparrow{
		Bird: Bird{Name: "Sparrow"},
	}

	penguin2 := Penguin{
		Bird: Bird{Name: "Penguin"},
	}

	// Sparrow can be used wherever FlyingBird is expected
	var flyingBird FlyingBird = sparrow2
	flyingBird.Fly()

	// Penguin is still a Bird
	penguin2.Eat()

	fmt.Println("Penguin does not implement FlyingBird because it cannot fly.")
}
