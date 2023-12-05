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
				simulateFastMouseMovementAndKeyPress()
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

func simulateFastMouseMovementAndKeyPress() {
	startTime := time.Now()
	for {
		// 获取当前时间
		currentTime := time.Now()

		// 每隔半秒快速旋转视角
		if currentTime.Sub(startTime) >= 500*time.Millisecond {
			// 获取当前鼠标位置
			x, y := robotgo.Location()

			// 快速水平移动鼠标
			newX := x + 50 // 增加移动的像素数
			robotgo.MoveSmooth(newX, y, 1.0, 10.0)

			// 重置开始时间
			startTime = currentTime
		}

		// 每隔三秒按下 Q 键和 E 键
		if currentTime.Sub(startTime) >= 3*time.Second {
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

			// 重置开始时间
			startTime = currentTime
		}

		// 控制循环速度
		time.Sleep(10 * time.Millisecond)
	}
}
