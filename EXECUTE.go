package main

import (
	"fmt"
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

var waiting bool = true

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

	go func() {
		for waiting {
			obj.GetInstruction(window, canvases, "Begin.jpg", &waiting)
		}
		fmt.Println("keepgoing")
	}()

	(*window).SetContent(containers.Stimuli)
	(*window).ShowAndRun()
}
