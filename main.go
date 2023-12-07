package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Neuvillette-Helper")
	label := widget.NewLabel("Hold down the middle mouse button to start spinning!")
	// 创建按钮
	startButton := widget.NewButton("Start", func() {
		// 这里添加启动项目的代码
		simulate()
	})

	//stopButton := widget.NewButton("Stop", func() {
	//	// 这里添加停止项目的代码
	//	os.Exit(1)
	//})

	// 将按钮添加到窗口
	myWindow.SetContent(container.NewVBox(
		label,
		startButton,
		//stopButton,
	))

	myWindow.ShowAndRun()
}
