package main

import (
	"fmt"
	"sync"
)

// Singleton
type Database struct {
	connection string
}

var (
	instance *Database
	once     sync.Once
)

// GetInstance returns the same Database instance every time.
// sync.Once guarantees that the initialization code runs exactly once, even when multiple goroutines call GetInstance() concurrently. It is simple, thread-safe, and idiomatic Go

// 1 way
// =================================================================
func GetInstance() *Database {

	once.Do(func() {
		fmt.Println("Creating database connection...")

		instance = &Database{
			connection: "DB Connection",
		}
	})

	return instance
}

// 2 way
// =================================================================
// var instance *Database

// func init() {
// 	instance = &Database{
// 		connection: "DB Connection",
// 	}
// }

// func GetInstance() *Database {
// 	return instance
// }

// 3 way
// =================================================================
// var (
// 	instance *Database
// 	mutex    sync.Mutex
// )

// func GetInstance() *Database {

// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	if instance == nil {
// 		instance = &Database{
// 			connection: "DB Connection",
// 		}
// 	}

// 	return instance
// }

// func main() {

// 	db1 := GetInstance()
// 	db2 := GetInstance()

// 	fmt.Println(db1.connection)
// 	fmt.Println(db2.connection)

// 	fmt.Println("Same instance:", db1 == db2)
// }
