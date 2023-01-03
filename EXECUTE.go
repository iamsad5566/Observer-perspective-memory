package main

import (
	"archive/zip"
	"fmt"
	"io"
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

var subject string
var width float32
var height float32
var group int
var showingTime int

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

	group, _ = strconv.Atoi(os.Getenv("GROUP"))
	showingTime, _ = strconv.Atoi(os.Getenv("SHOWING_TIME"))
	subject = os.Getenv("SUBJECT_NUM")
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
		// I1
		obj.GetInstruction(window, canvases, instructFile.Instructions[0], &waiting)
		waitKeyPress()

		// I2
		obj.GetInstruction(window, canvases, instructFile.Instructions[1], &waiting)
		waitKeyPress()

		// I3
		obj.GetInstruction(window, canvases, instructFile.Instructions[2], &waiting)
		waitKeyPress()

		// I4
		obj.GetInstruction(window, canvases, instructFile.Instructions[3], &waiting)
		waitKeyPress()

		// I5
		event.CaptureEscape(window)
		obj.GetPreTrainStimulus(window, canvases, instructFile.Instructions[4])
		time.Sleep(time.Second * 5)

		// I6
		correct := false
		obj.PreTrainInstruction(window, canvases, instructFile.Instructions[5], &correct, &waiting)
		waitKeyPress()

		if correct {
			// I7
			obj.GetInstruction(window, canvases, instructFile.Instructions[6], &waiting)
			waitKeyPress()
		} else {
			// I8
			obj.GetInstruction(window, canvases, instructFile.Instructions[7], &waiting)
			waitKeyPress()
		}

		// I9
		event.CaptureEscape(window)
		obj.GetPreTrainStimulus(window, canvases, instructFile.Instructions[8])
		time.Sleep(time.Second * 5)

		// I10
		correct = false
		obj.PreTrainInstruction(window, canvases, instructFile.Instructions[9], &correct, &waiting)
		waitKeyPress()

		if correct {
			// I11
			obj.GetInstruction(window, canvases, instructFile.Instructions[10], &waiting)
			waitKeyPress()
		} else {
			// I12
			obj.GetInstruction(window, canvases, instructFile.Instructions[11], &waiting)
			waitKeyPress()
		}

		// I13
		obj.GetInstruction(window, canvases, instructFile.Instructions[12], &waiting)
		waitKeyPress()

		// Showing pics
		content.RemoveAll()
		content.Add(containers.Stimuli)
		content.Refresh()
		event.CaptureEscape(window)

		for _, str := range pictureFile.Slice {
			obj.GetStimulus(canvases, str)
			time.Sleep(time.Second * time.Duration(showingTime))
			canvases.Picture.File = pictureFile.Mask
			canvases.Picture.Refresh()
			time.Sleep(time.Second)
		}

		// Response phase
		content.RemoveAll()
		content.Add(containers.Instruction)
		content.Refresh()

		// I14
		if group == 1 {
			obj.GetInstruction(window, canvases, instructFile.Instructions[14], &waiting)
			waitKeyPress()
		} else {
			// I15
			obj.GetInstruction(window, canvases, instructFile.Instructions[14], &waiting)
			waitKeyPress()
		}

		// Show pictures
		// Response phase
		content.RemoveAll()
		content.Add(containers.Stimuli)
		content.Refresh()

		for i, str := range pictureFile.Slice {
			obj.GetResponseToStimulus(window, canvases, str, &waiting, &result, i)
			waitKeyPress()
			canvases.Picture.File = pictureFile.Mask
			canvases.Picture.Refresh()
			time.Sleep(time.Second)
		}

		// I16
		content.RemoveAll()
		content.Add(containers.Instruction)
		content.Refresh()

		obj.END(window, canvases, instructFile.Instructions[15], &waiting)
		exportData()
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

func exportData() {
	file, _ := os.Create("result/" + subject + ".csv")
	file.WriteString("trial,ratio,group\n")
	for i, ratio := range result {
		file.WriteString(fmt.Sprintf("%d,%f,%d\n", i+1, ratio, group))
	}
	file.Close()

	zipFile, _ := os.Create("result/" + subject + ".zip")
	defer zipFile.Close()
	writer := zip.NewWriter(zipFile)

	file, _ = os.Open("result/" + subject + ".csv")
	w, _ := writer.Create(subject + ".csv")
	io.Copy(w, file)
	writer.Close()
}
