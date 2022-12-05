package main

import (
	"observerPerspective/event"
	"observerPerspective/win"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var width float32 = 1920
var height float32 = 1440

const (
	ShowingPhase = iota
	ResponsePhase
)

func main() {
	openGUI()
}

// openGUI starts a new program
func openGUI() {
	guiApp := app.New()
	guiApp.Settings().SetTheme(&myTheme{})
	window := guiApp.NewWindow("Observer perspective memory")
	window.Resize(fyne.NewSize(width, height))
	window.SetFixedSize(false)
	window.SetFullScreen(true)
	window.SetContent(procedureController(window))
	window.ShowAndRun()
}

func procedureController(window fyne.Window) fyne.CanvasObject {

	go event.CaptureEscape(window)

	return win.BuildStimuliWin(0, 0)
}
