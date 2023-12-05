package main

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"time"
)

var simulating = false

func main() {
	// Register a mouse event for the middle button
	hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["center"] {
			simulating = !simulating
			if simulating {
				log.Default().Println("Simulating starts")
				// Start mouse movement and key press simulations in separate goroutines
				go simulateMouseMovement()
				go simulateKeyPress()
			} else {
				log.Default().Println("Simulating ends")
			}
		}
	})

	// Start listening for mouse and keyboard events
	s := hook.Start()
	defer hook.End()

	// Process events
	<-hook.Process(s)
}

func simulateMouseMovement() {
	err := robotgo.MouseDown("left")
	// Press the left mouse button
	if err != nil {
		return
	}
	log.Default().Println("Mouse left button pressed")
	for simulating {
		// Get the current mouse position
		x, y := robotgo.Location()

		// Move the mouse horizontally while keeping the vertical position unchanged
		newX := x + 500 // Move 500 pixels each time
		robotgo.Move(newX, y)

		// Control the speed of the loop
		time.Sleep(10 * time.Millisecond)
	}
	err = robotgo.MouseUp("left")
	// Release the left mouse button
	if err != nil {
		return
	}
	log.Default().Println("Mouse left button released")
}

func simulateKeyPress() {
	for simulating {
		err := robotgo.KeyTap("e")
		if err != nil {
			log.Default().Println("Error pressing E:", err)
			return
		}
		log.Default().Println("E pressed")

		time.Sleep(100 * time.Millisecond) // Short delay

		err = robotgo.KeyTap("q")
		if err != nil {
			log.Default().Println("Error pressing Q:", err)
			return
		}
		log.Default().Println("Q pressed")

		time.Sleep(3 * time.Second)
		// Press Q and E keys every three seconds
	}
}
