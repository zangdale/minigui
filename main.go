package main

import (
	"minigui/server"

	"github.com/getlantern/systray"
)

const (
	version = "v2021.1.30"
)

func main() {
	go server.Init()
	systray.Run(onReady("https://www.baidu.com"), onExit)
}
