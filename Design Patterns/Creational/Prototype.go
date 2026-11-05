// **Type assertion**
// is used to tell Go that an interface value actually contains a specific concrete type. In our Prototype example, `Clone()` returns the `Enemy` interface, so Go only knows that the returned object is an `Enemy`; it doesn't know about fields like `name`, `health`, or `damage`. When we write `baseBoss.Clone().(*BossEnemy)`, we are telling Go that this `Enemy` is actually a `*BossEnemy`, so we can access and modify its fields, such as `boss1.name` and `boss1.damage`.

package main

import "fmt"

// Prototype
type Enemy interface {
	Clone() Enemy
	Show()
}

// Concrete Prototype
type BossEnemy struct {
	name     string
	health   int
	damage   int
	gpuModel string
	aiModel  string
}

// Clone creates a copy of an already prepared enemy.
func (b *BossEnemy) Clone() Enemy {
	return &BossEnemy{
		name:     b.name,
		health:   b.health,
		damage:   b.damage,
		gpuModel: b.gpuModel,
		aiModel:  b.aiModel,
	}
}

func (b *BossEnemy) Show() {
	fmt.Println(
		b.name,
		"| Health:", b.health,
		"| Damage:", b.damage,
		"| GPU Model:", b.gpuModel,
		"| AI Model:", b.aiModel,
	)
}

func main() {

	// Creating this boss is expensive.
	//
	// Imagine this involves:
	// - Loading a huge 3D model into GPU memory
	// - Loading textures
	// - Loading animations
	// - Loading AI model
	// - Running expensive GPU processing
	//
	// We don't want to repeat all this work for every enemy.
	baseBoss := &BossEnemy{
		name:     "Boss",
		health:   1000,
		damage:   200,
		gpuModel: "Huge-Dragon-3D-Model",
		aiModel:  "Advanced-Boss-AI",
	}

	// Clone the already prepared boss.
	boss1 := baseBoss.Clone().(*BossEnemy)

	// Tweak only what is different.
	boss1.name = "Fire Dragon"
	boss1.damage = 300

	// Clone again.
	boss2 := baseBoss.Clone().(*BossEnemy)

	// Tweak only what is different.
	boss2.name = "Ice Dragon"
	boss2.health = 1500

	boss1.Show()
	boss2.Show()
}
