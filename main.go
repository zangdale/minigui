package main

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const (
	version = "v2021.1.30"
)

var (
	mainApp          = app.NewWithID("com.minigui") // app 的 main 信息
	logWindow        fyne.Window
	logEntry         *widget.Entry
	indexWindow      fyne.Window
	trayIconResource []byte
)

func main() {
	setupIcon()                     // 设置页面图标
	go startTray()                  // 状态栏图标
	logWindow = buildLogWindow()    // 日志页面
	indexWindow = buildPrefWindow() // 主页面
	go start()                      // 开始工作
	mainApp.Run()
}

func setupIcon() {
	mainApp.SetIcon(pngIconResource)
	if runtime.GOOS == "windows" {
		trayIconResource = icoIconResource.StaticContent
	} else {
		trayIconResource = pngIconResource.StaticContent
	}
	// 主题
	mainApp.Settings().SetTheme(DefaultTheme())
}
