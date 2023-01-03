package main

import (
	"fmt"
	"log"
	"observerPerspective/event"
	"observerPerspective/material"
	"observerPerspective/obj"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/joho/godotenv"
)

var width float32
var height float32

var waiting bool = true
var instructFile *material.InstructFile = &material.InstructFile{}
var pictureFile *material.PictureFile = &material.PictureFile{}
var canvases *obj.Canvases = &obj.Canvases{}
var containers *obj.Containers = &obj.Containers{}

// Result
var result = []float32{}

func main() {
	loadEnv()
	openGUI()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	strWidth := os.Getenv("WIDTH")
	strHeight := os.Getenv("HEIGHT")
	width64, _ := strconv.ParseFloat(strWidth, 32)
	height64, _ := strconv.ParseFloat(strHeight, 32)
	width = float32(width64)
	height = float32(height64)
}

// openGUI starts a new program
func openGUI() {
	guiApp := app.New()
	intializeObjects()
	guiApp.Settings().SetTheme(&myTheme{})
	window := guiApp.NewWindow("Observer perspective memory")
	window.Resize(fyne.NewSize(width, height))
	window.SetFixedSize(false)
	window.SetFullScreen(true)

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
			obj.GetStimulus(canvases, str)
			time.Sleep(time.Second * 5)
			canvases.Picture.File = pictureFile.Mask
			canvases.Picture.Refresh()
			time.Sleep(time.Second)
		}

		// Response phase
		content.RemoveAll()
		content.Add(containers.Instruction)
		content.Refresh()

		obj.GetInstruction(window, canvases, instructFile.Prepare, &waiting)
		waitKeyPress()

		// Show pictures
		// Response phase
		content.RemoveAll()
		content.Add(containers.Stimuli)
		content.Refresh()

		for _, str := range pictureFile.Slice {
			obj.GetResponseToStimulus(window, canvases, str, &waiting, &result)
			waitKeyPress()
			canvases.Picture.File = pictureFile.Mask
			canvases.Picture.Refresh()
			time.Sleep(time.Second)
		}

		fmt.Println(result)
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
