package main

import "fmt"

// Subsystem 1
type TV struct{}

func (t *TV) On() {
	fmt.Println("TV ON")
}

func (t *TV) Off() {
	fmt.Println("TV OFF")
}

// Subsystem 2
type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Sound System ON")
}

func (s *SoundSystem) Off() {
	fmt.Println("Sound System OFF")
}

// Subsystem 3
type Lights struct{}

func (l *Lights) Dim() {
	fmt.Println("Lights dimmed")
}

func (l *Lights) On() {
	fmt.Println("Lights ON")
}

// Facade
type HomeTheaterFacade struct {
	tv     *TV
	sound  *SoundSystem
	lights *Lights
}

func (h *HomeTheaterFacade) WatchMovie() {
	fmt.Println("Preparing to watch movie...")

	h.lights.Dim()
	h.tv.On()
	h.sound.On()

	fmt.Println("Movie started")
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Stopping movie...")

	h.sound.Off()
	h.tv.Off()
	h.lights.On()

	fmt.Println("Movie ended")
}

func main() {

	tv := &TV{}
	sound := &SoundSystem{}
	lights := &Lights{}

	homeTheater := &HomeTheaterFacade{
		tv:     tv,
		sound:  sound,
		lights: lights,
	}

	// Client only interacts with Facade
	homeTheater.WatchMovie()

	fmt.Println()

	homeTheater.EndMovie()
}
