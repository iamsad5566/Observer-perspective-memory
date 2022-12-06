package main

import (
	"observerPerspective/event"
	"observerPerspective/material"
	"observerPerspective/obj"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var width float32 = 1920
var height float32 = 1440

const (
	InstructionPhase = iota
	ShowingPhase
	ResponsePhase
)

var waiting bool = true
var instructFile *material.InstructFile = &material.InstructFile{}
var pictureFile *material.PictureFile = &material.PictureFile{}
var canvases *obj.Canvases = &obj.Canvases{}
var containers *obj.Containers = &obj.Containers{}

func main() {
	openGUI()
}

// openGUI starts a new program
func openGUI() {
	guiApp := app.New()
	intializeObjects()
	guiApp.Settings().SetTheme(&myTheme{})
	window := guiApp.NewWindow("Observer perspective memory")
	window.Resize(fyne.NewSize(width, height))
	window.SetFixedSize(false)
	window.SetFullScreen(false)
	procedureController(&window)
}

func procedureController(window *fyne.Window) {
	content := container.NewCenter(containers.Instruction)
	go func() {
		// Fist Instruction
		obj.GetInstruction(window, canvases, instructFile.Begin, &waiting)
		waitKeyPress()

		// Description
		obj.GetInstruction(window, canvases, instructFile.Description, &waiting)
		waitKeyPress()

		// Prepare
		obj.GetInstruction(window, canvases, instructFile.Prepare, &waiting)
		waitKeyPress()

		// Showing pics
		content.RemoveAll()
		content.Add(containers.Stimuli)
		content.Refresh()
		event.CaptureEscape(window)

		for _, str := range pictureFile.Slice {
			obj.GetStimulus(window, canvases, str)
			time.Sleep(time.Second * 1)
			canvases.Picture.File = pictureFile.Mask
			canvases.Picture.Refresh()
			time.Sleep(time.Second)
		}

		// Response pahse
		content.RemoveAll()
		content.Add(containers.Instruction)
		content.Refresh()

		obj.GetInstruction(window, canvases, instructFile.Prepare, &waiting)
		waitKeyPress()

		// Shuffle first
		pictureFile.ShuffleSlice()

		// Show pictures
		// Response pahse
		content.RemoveAll()
		content.Add(containers.Stimuli)
		content.Refresh()

		for _, str := range pictureFile.Slice {
			obj.GetResponseToStimulus(window, canvases, str, &waiting)
			waitKeyPress()
			canvases.Picture.File = pictureFile.Mask
			canvases.Picture.Refresh()
			time.Sleep(time.Second)
		}
	}()

	(*window).SetContent(content)
	(*window).ShowAndRun()
}

func intializeObjects() {
	instructFile.Load()
	pictureFile.Load()
	canvases.Load(instructFile, pictureFile)
	containers.Load(canvases)
}

func waitKeyPress() {
	for waiting {
		continue
	}
	waiting = true
}
