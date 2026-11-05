// It lets us treat individual objects and groups of objects in the same way
package main

import "fmt"

// Component
type FileSystem interface {
	Show()
}

// Leaf
type File struct {
	name string
}

func (f *File) Show() {
	fmt.Println("File:", f.name)
}

// Composite
type Folder struct {
	name     string
	children []FileSystem
}

func (f *Folder) Add(component FileSystem) {
	f.children = append(f.children, component)
}

func (f *Folder) Show() {
	fmt.Println("Folder:", f.name)

	for _, child := range f.children {
		child.Show()
	}
}

// func main() {

// 	// Files
// 	file1 := &File{name: "resume.pdf"}
// 	file2 := &File{name: "photo.jpg"}
// 	file3 := &File{name: "notes.txt"}

// 	// Folders
// 	documents := &Folder{name: "Documents"}
// 	pictures := &Folder{name: "Pictures"}

// 	// Add files to folders
// 	documents.Add(file1)
// 	documents.Add(file3)

// 	pictures.Add(file2)

// 	// Root folder
// 	root := &Folder{name: "Home"}

// 	root.Add(documents)
// 	root.Add(pictures)

// 	// Show everything
// 	root.Show()
// }
