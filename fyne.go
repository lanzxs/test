package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()            // 创建一个窗口
	w := a.NewWindow("Hello") // 给窗口设置一个标题

	hello1 := widget.NewLabel("111111111111")  // 创建一个标签
	hello2 := widget.NewLabel("222222222222")  // 创建一个标签
	hello3 := widget.NewLabel("333333333333")  // 创建一个标签
	button := widget.NewButton("Hi!", func() { // 按钮控件  按钮是带点击事件的
		hello1.SetText("Welcome :)") // 修改标签的 标题
	})
	// Grid 布局--- 没整明白是什么布局
	//w.SetContent(container.NewGridWithColumns(2, hello1, hello2, hello3, button, button, button))
	w.SetContent(container.NewAdaptiveGrid(2, hello1, hello2, hello3, button))
	//w.Resize(fyne.NewSize(500, 500)) // 设置窗口大小
	w.ShowAndRun()
}
