package main

import (
	"fmt"
	"observerPerspective/event"
	"observerPerspective/material"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var width float32 = 1600
var height float32 = 1200

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
	// window.SetFullScreen(true)
	window.SetContent(test(window))
	window.ShowAndRun()
}

func test(window fyne.Window) fyne.CanvasObject {
	stimuli := material.Stimuli{}
	stimuli.Load()
	con := canvas.NewImageFromFile(stimuli.Array[0])
	con.SetMinSize(fyne.Size{Width: 300, Height: 400})
	fmt.Println(con.MinSize())
	con.FillMode = canvas.ImageFillContain
	content := container.NewCenter(con)

	go event.CaptureEscape(window)

	return content
}
