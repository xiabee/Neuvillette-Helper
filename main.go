package main

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"time"
)

func main() {
	fmt.Println("Listening for mouse left button hold...")

	var leftButtonHeld bool

	// 注册鼠标左键按下事件
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			leftButtonHeld = true
			fmt.Println("Left mouse button pressed")
		}
	})

	// 注册鼠标左键释放事件
	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			leftButtonHeld = false
			fmt.Println("Left mouse button released")
		}
	})

	// 启动事件监听
	s := hook.Start()
	<-hook.Process(s)

	for {
		if leftButtonHeld {
			fmt.Println("Left mouse button is being held...")
			time.Sleep(500 * time.Millisecond) // 每半秒检查一次
		} else {
			time.Sleep(100 * time.Millisecond) // 减少循环频率以减少CPU使用
		}
	}
}
