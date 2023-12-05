package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"time"
)

func main() {
	var clicked = false
	hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["center"] {
			fmt.Printf("mouse center @ %v - %v\n", e.X, e.Y)
			clicked = !clicked
			if clicked {
				simulateMouseMovementAndKeyPress()
			}
		} else if e.Button == hook.MouseMap["right"] {
			hook.End()
			return
		}
	})
	s := hook.Start()
	defer hook.End()

	<-hook.Process(s)
}

func simulateMouseMovementAndKeyPress() {
	for {
		// 获取当前鼠标位置
		x, y := robotgo.Location()

		// 快速水平移动鼠标，保持垂直位置不变
		newX := x + 10 // 每次移动增加10像素
		robotgo.Move(newX, y)

		// 每隔三秒按下 Q 键和 E 键
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
			// 控制循环速度，以允许响应用户的鼠标移动
			time.Sleep(10 * time.Millisecond)
		}
	}
}
