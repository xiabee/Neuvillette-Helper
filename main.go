package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"time"
)

/*
#include <windows.h>

void simulateMouseMove(int x, int y) {
    mouse_event(MOUSEEVENTF_MOVE, x, y, 0, 0);
}
*/
import "C"

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
				// go simulateKeyPress()
			} else {
				log.Default().Println("Simulating ends")
			}
		}

	}
	// Register a mouse event for the middle button
	hook.Register(hook.MouseHold, []string{}, handleMouseEvent)
	hook.Register(hook.MouseDown, []string{}, handleMouseEvent)
	// hook.Register(hook.MouseUp, []string{}, handleMouseEvent)

	// Start listening for mouse and keyboard events
	s := hook.Start()
	defer hook.End()

	<-hook.Process(s)
}

func simulateMouseMovement() {
	err := robotgo.MouseDown("left")
	// Press the left mouse buttonn
	if err != nil {
		return
	}
	log.Default().Println("Mouse left button pressed")
	for simulating {
		//// 使用 CGO 调用移动鼠标
		C.simulateMouseMove(1000, 0)
		// 控制循环速度
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
