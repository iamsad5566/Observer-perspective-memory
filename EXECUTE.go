package main

import (
	"observerPerspective/event"
	"observerPerspective/obj"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var width float32 = 1920
var height float32 = 1440

const (
	InstructionPhase = iota
	ShowingPhase
	ResponsePhase
)

var keepGoing bool = false

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
	window.SetFullScreen(false)
	procedureController(&window)
}

func procedureController(window *fyne.Window) {
	canvases := &obj.Canvases{}
	canvases.Load()
	containers := &obj.Containers{}
	containers.Load(canvases)
	// event.CaptureEscape(window)
	event.CaptureZoom(window, ResponsePhase, canvases.Picture)
	obj.GetInstruction(window, canvases)
	(*window).SetContent(containers.Stimuli)
	(*window).ShowAndRun()
}
