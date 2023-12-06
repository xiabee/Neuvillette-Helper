package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"time"
)

var simulating = false

func main() {
	fmt.Println("===========|| Press the middle mouse button to start the simulation ||===========")
	handleMouseEvent := func(e hook.Event) {
		fmt.Printf("Mouse event triggered: %v, Button: %d\n", e.Kind, e.Button)
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

	}
	// Register a mouse event for the middle button
	hook.Register(hook.MouseHold, []string{}, handleMouseEvent)
	hook.Register(hook.MouseDown, []string{}, handleMouseEvent)
	// eqeqhook.Register(hook.MouseUp, []string{}, handleMouseEvent)

	// Start listening for mouse and keyboard events
	s := hook.Start()
	defer hook.End()

	// Process events
	<-hook.Process(s)
	//base()
}

func simulateMouseMovement() {
	err := robotgo.MouseDown("left")
	// Press the left mouse button
	if err != nil {
		return
	}
	log.Default().Println("Mouse left button pressed")
	for simulating {
		//Get the current mouse position
		x, y := robotgo.Location()

		// Move the mouse horizontally while keeping the vertical position unchanged
		newX := x + 100 // Move 100 pixels each time
		robotgo.Move(newX, y/2)

		// Control the speed of the loop
		time.Sleep(10 * time.Millisecond)
	}
	err = robotgo.MouseUp("left")
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

func base() {
	fmt.Println("hook start...")
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
		if ev.Keychar == 'q' {
			break
		}
	}
}
