package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {
	//hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
	//
	//})
	CheckMouseMiddle()
}

func CheckMouseMiddle() {
	hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["center"] {
			fmt.Printf("mouse center @ %v - %v\n", e.X, e.Y)
		} else if e.Button == hook.MouseMap["right"] {
			hook.End()
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
