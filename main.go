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
				simulateMouseAndKeyboard()
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

func move() {
	robotgo.MouseSleep = 100
	robotgo.Move(100, 200)
	robotgo.MoveRelative(10, -200)

	// move the mouse to 100, 200
	robotgo.Move(100, 200)

	// drag mouse with smooth
	robotgo.DragSmooth(10, 10)
	robotgo.DragSmooth(100, 200, 1.0, 100.0)

	// smooth move the mouse to 100, 200
	robotgo.MoveSmooth(100, 200)
	robotgo.MoveSmooth(100, 200, 1.0, 100.0)
	robotgo.MoveSmoothRelative(10, -100, 1.0, 30.0)

	for i := 0; i < 1080; i += 1000 {
		fmt.Println("i: ", i)
		// MoveMouse(800, i)
		robotgo.Move(800, i)
	}
}

func click() {

	// click the left mouse button
	robotgo.Click()

	// click the right mouse button
	robotgo.Click("right", false)

	// double-click the left mouse button
	robotgo.Click("left", true)
}

func get() {
	// gets the mouse coordinates
	x, y := robotgo.Location()
	fmt.Println("pos:", x, y)
	if x == 456 && y == 586 {
		fmt.Println("mouse...", "586")
	}

	robotgo.Move(x, y)
}

func toggleAndScroll() {
	// scrolls the mouse either up
	robotgo.ScrollDir(10, "up")
	robotgo.ScrollDir(10, "right")

	robotgo.Scroll(100, 10)
	robotgo.Scroll(0, -10)

	robotgo.Toggle("left")
	robotgo.Toggle("left", "up")

	// toggles the right mouse button
	robotgo.Toggle("right")
	robotgo.Toggle("right", "up")
}

func simulateMouseAndKeyboard() {

	robotgo.MouseSleep = 100
	for {
		// 模拟鼠标快速水平移动
		x, y := robotgo.Location()        // 获取当前鼠标位置
		robotgo.Move(x+10, y)             // 向右移动10像素
		time.Sleep(10 * time.Millisecond) // 控制移动速度

		// 每隔三秒按下 Q 键和 E 键
		time.Sleep(3 * time.Second)
		err := robotgo.KeyTap("q")
		if err != nil {
			return
		}
		log.Default().Println("Q pressed")

		time.Sleep(100 * time.Millisecond) // 短暂延迟
		err = robotgo.KeyTap("e")
		if err != nil {
			return
		}
		log.Default().Println("E pressed")
	}
}
