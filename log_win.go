package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func buildLogWindow() fyne.Window {
	logWindow := mainApp.NewWindow("Logs")
	logWindow.SetIcon(pngIconResource)
	logWindow.CenterOnScreen()
	logWindow.Resize(fyne.NewSize(600, 800))
	logWindow.SetCloseIntercept(func() {
		logWindow.Hide()
	})

	logEntry = widget.NewMultiLineEntry()
	logEntry.Disable()

	logWindow.SetContent(container.NewScroll(logEntry))

	logWindow.Hide()
	return logWindow
}

func newLogError(text string, err error) {
	logEntry.Text += time.Now().Format(time.RFC3339) + "\tERROR\t" + text + "\t" + err.Error() + "\n"
}

func newLogInfo(text string) {
	logEntry.Text += time.Now().Format(time.RFC3339) + "\tINFO\t" + text + "\n"
}
