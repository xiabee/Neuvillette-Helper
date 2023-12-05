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
	hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["center"] {
			fmt.Printf("mouse center @ %v - %v\n", e.X, e.Y)
			simulating = !simulating
			if simulating {
				go simulateMouseMovementAndKeyPress()
			}
		}
	})

	s := hook.Start()
	defer hook.End()

	<-hook.Process(s)
}

func simulateMouseMovementAndKeyPress() {
	for simulating {
		// Get the current mouse position
		x, y := robotgo.Location()

		// Move the mouse quickly horizontally, keeping the vertical position unchanged
		newX := x + 10 // Each move adds 10 pixels
		robotgo.Move(newX, y)

		// Press Q and E every three seconds
		select {
		case <-time.After(3 * time.Second):
			err := robotgo.KeyTap("e")
			if err != nil {
				return
			}
			log.Default().Println("E pressed")
			time.Sleep(100 * time.Millisecond) // 短暂延迟
			err = robotgo.KeyTap("q")
			if err != nil {
				return
			}
			log.Default().Println("Q pressed")
		default:
			// Control loop speed to allow response to user's mouse movements
			time.Sleep(10 * time.Millisecond)
		}
	}
}
