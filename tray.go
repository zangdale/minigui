package main

import (
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

func onReady() {
	systray.SetIcon(trayIconResource)
	systray.SetTitle("My APP")
	systray.SetTooltip("My APP")

	mPref := systray.AddMenuItem("Main APP", "Open Main APP")
	mLog := systray.AddMenuItem("Logs", "See Logs")
	mQuit := systray.AddMenuItem("Quit", "Quit Go Reddit WallPaper")
	systray.AddSeparator()
	mAboutLink := systray.AddMenuItem("About", "About Me!")

	for {
		select {
		case <-mAboutLink.ClickedCh:
			open.Run("https://www.github.com/getbuguai")
		case <-mPref.ClickedCh:
			indexWindow.Show()

		case <-mLog.ClickedCh:
			logWindow.Show()

		case <-mQuit.ClickedCh:
			systray.Quit()
			return
		}
	}
}
func onExit() {
	mainApp.Quit()
}

func startTray() {
	systray.Run(onReady, onExit)
}
