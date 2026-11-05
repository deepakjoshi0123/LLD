package main

import "fmt"

// Subject
type VideoService interface {
	PlayVideo()
}

// Real Subject
type RealVideoService struct{}

func (r *RealVideoService) PlayVideo() {
	fmt.Println("Playing video...")
}

// Proxy
type VideoServiceProxy struct {
	service *RealVideoService
}

func (p *VideoServiceProxy) PlayVideo() {

	fmt.Println("Checking user access...")

	// Access check
	allowed := true

	if !allowed {
		fmt.Println("Access denied")
		return
	}

	// Forward request to real object
	p.service.PlayVideo()
}

// Client
// func main() {

// 	realService := &RealVideoService{}

// 	proxy := &VideoServiceProxy{
// 		service: realService,
// 	}

// 	// Client uses Proxy instead of RealVideoService
// 	proxy.PlayVideo()
// }