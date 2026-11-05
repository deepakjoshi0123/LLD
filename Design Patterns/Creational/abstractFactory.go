package main

import "fmt"

// --------------------
// Product interfaces
// --------------------

// Button
type Button interface {
	Render()
}

// Checkbox
type Checkbox interface {
	Render()
}

// --------------------
// Windows Products
// --------------------

type WindowsButton struct{}

func (w *WindowsButton) Render() {
	fmt.Println("Rendering Windows Button")
}

type WindowsCheckbox struct{}

func (w *WindowsCheckbox) Render() {
	fmt.Println("Rendering Windows Checkbox")
}

// --------------------
// Mac Products
// --------------------

type MacButton struct{}

func (m *MacButton) Render() {
	fmt.Println("Rendering Mac Button")
}

type MacCheckbox struct{}

func (m *MacCheckbox) Render() {
	fmt.Println("Rendering Mac Checkbox")
}

// --------------------
// Abstract Factory
// --------------------

type UIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// --------------------
// Windows Factory
// --------------------

type WindowsFactory struct{}

func (w *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (w *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

// --------------------
// Mac Factory
// --------------------

type MacFactory struct{}

func (m *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (m *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

// --------------------
// Client
// --------------------

func RenderUI(factory UIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	button.Render()
	checkbox.Render()
}

// func main() {

// 	// Windows UI
// 	windowsFactory := &WindowsFactory{}
// 	RenderUI(windowsFactory)

// 	fmt.Println()

// 	// Mac UI
// 	macFactory := &MacFactory{}
// 	RenderUI(macFactory)
// }

//``````````````````````````````````````````````

//with simple factory 

// package main

// import "fmt"

// // --------------------
// // Product interfaces
// // --------------------

// type Button interface {
// 	Render()
// }

// type Checkbox interface {
// 	Render()
// }

// // --------------------
// // Windows Products
// // --------------------

// type WindowsButton struct{}

// func (w *WindowsButton) Render() {
// 	fmt.Println("Rendering Windows Button")
// }

// type WindowsCheckbox struct{}

// func (w *WindowsCheckbox) Render() {
// 	fmt.Println("Rendering Windows Checkbox")
// }

// // --------------------
// // Mac Products
// // --------------------

// type MacButton struct{}

// func (m *MacButton) Render() {
// 	fmt.Println("Rendering Mac Button")
// }

// type MacCheckbox struct{}

// func (m *MacCheckbox) Render() {
// 	fmt.Println("Rendering Mac Checkbox")
// }

// // --------------------
// // Simple Factory
// // --------------------

// func CreateButton(os string) Button {

// 	switch os {
// 	case "windows":
// 		return &WindowsButton{}

// 	case "mac":
// 		return &MacButton{}

// 	default:
// 		return nil
// 	}
// }

// func CreateCheckbox(os string) Checkbox {

// 	switch os {
// 	case "windows":
// 		return &WindowsCheckbox{}

// 	case "mac":
// 		return &MacCheckbox{}

// 	default:
// 		return nil
// 	}
// }

// // --------------------
// // Client
// // --------------------

// func main() {

// 	os := "windows"

// 	button := CreateButton(os)
// 	checkbox := CreateCheckbox(os)

// 	button.Render()
// 	checkbox.Render()
// }