package main

import "fmt"

// Observer
type Subscriber interface {
	Update(video string)
}

// Subject
type YouTubeChannel struct {
	subscribers []Subscriber
}

// Add subscriber
func (y *YouTubeChannel) Subscribe(s Subscriber) {
	y.subscribers = append(y.subscribers, s)
}

// Remove subscriber
func (y *YouTubeChannel) Unsubscribe(s Subscriber) {
	for i, subscriber := range y.subscribers {
		if subscriber == s {
			y.subscribers = append(y.subscribers[:i], y.subscribers[i+1:]...)
			return
		}
	}
}

// Notify all subscribers
func (y *YouTubeChannel) UploadVideo(video string) {
	fmt.Println("New video uploaded:", video)

	for _, subscriber := range y.subscribers {
		subscriber.Update(video)
	}
}

// Concrete Observer
type User struct {
	name string
}

func (u *User) Update(video string) {
	fmt.Println(u.name, "received notification for:", video)
}

// func main() {

// 	// Create channel
// 	channel := &YouTubeChannel{}

// 	// Create subscribers
// 	user1 := &User{name: "Rahul"}
// 	user2 := &User{name: "Amit"}
// 	user3 := &User{name: "Deepak"}

// 	// Subscribe users
// 	channel.Subscribe(user1)
// 	channel.Subscribe(user2)
// 	channel.Subscribe(user3)

// 	// Upload video
// 	channel.UploadVideo("Go Observer Pattern")

// 	// Unsubscribe one user
// 	channel.Unsubscribe(user2)

// 	fmt.Println()

// 	// Upload another video
// 	channel.UploadVideo("Go Design Patterns")
// }
