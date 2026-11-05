package main

// Component
type Coffee interface {
	Cost() int
	Description() string
}

// Concrete Component
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() int {
	return 100
}

func (c *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Decorator
type MilkDecorator struct {
	coffee Coffee
}

func (m *MilkDecorator) Cost() int {
	return m.coffee.Cost() + 30
}

func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + " + Milk"
}

// Another Decorator
type SugarDecorator struct {
	coffee Coffee
}

func (s *SugarDecorator) Cost() int {
	return s.coffee.Cost() + 10
}

func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + " + Sugar"
}

// Another Decorator
type WhippedCreamDecorator struct {
	coffee Coffee
}

func (w *WhippedCreamDecorator) Cost() int {
	return w.coffee.Cost() + 40
}

func (w *WhippedCreamDecorator) Description() string {
	return w.coffee.Description() + " + Whipped Cream"
}

// func main() {

// 	coffee := &SimpleCoffee{}

// 	// Add Milk
// 	coffeeWithMilk := &MilkDecorator{
// 		coffee: coffee,
// 	}

// 	// Add Sugar
// 	coffeeWithMilkAndSugar := &SugarDecorator{
// 		coffee: coffeeWithMilk,
// 	}

// 	// Add Whipped Cream
// 	finalCoffee := &WhippedCreamDecorator{
// 		coffee: coffeeWithMilkAndSugar,
// 	}

// 	fmt.Println(finalCoffee.Description())
// 	fmt.Println("Cost:", finalCoffee.Cost())
// }
