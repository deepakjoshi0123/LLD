package main

import "fmt"

// Product
type Computer struct {
	CPU     string
	RAM     int
	Storage int
	GPU     string
	WiFi    bool
}

func (c *Computer) Show() {
	fmt.Println("CPU:", c.CPU)
	fmt.Println("RAM:", c.RAM, "GB")
	fmt.Println("Storage:", c.Storage, "GB")
	fmt.Println("GPU:", c.GPU)
	fmt.Println("WiFi:", c.WiFi)
}

// Builder
type ComputerBuilder struct {
	computer *Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{
		computer: &Computer{},
	}
}

func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ComputerBuilder) SetRAM(ram int) *ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *ComputerBuilder) SetStorage(storage int) *ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *ComputerBuilder) SetGPU(gpu string) *ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *ComputerBuilder) SetWiFi(wifi bool) *ComputerBuilder {
	b.computer.WiFi = wifi
	return b
}

// Build final object
func (b *ComputerBuilder) Build() *Computer {
	return b.computer
}

// func main() {

// 	computer := NewComputerBuilder().
// 		SetCPU("Intel i7").
// 		SetRAM(32).
// 		SetStorage(1000).
// 		SetGPU("NVIDIA RTX 5070").
// 		SetWiFi(true).
// 		Build()

// 	computer.Show()
// }
