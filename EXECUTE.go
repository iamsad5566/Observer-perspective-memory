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
	preTrialCorrection := []bool{}

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

		// I5 -> FP 1~6
		content.RemoveAll()
		content.Add(containers.Stimuli)
		content.Refresh()
		event.CaptureEscape(window)
		for i := 1; i <= 5; i++ {
			obj.GetPreTrainStimulus(canvases, pictureFile.FPSlice[i])
			time.Sleep(time.Second * time.Duration(showingTime))
			canvases.Picture.File = pictureFile.Mask
			canvases.Picture.Refresh()
			time.Sleep(time.Second)
			content.RemoveAll()
			content.Add(containers.Instruction)
			content.Refresh()
			correct := false
			obj.PreTrainInstruction(window, canvases, instructFile.Instructions[5], &correct, &waiting)
			waitKeyPress()
			preTrialCorrection = append(preTrialCorrection, correct)
			if correct {
				// I7
				obj.GetInstruction(window, canvases, instructFile.Instructions[6], &waiting)
				waitKeyPress()
			} else {
				// I8
				obj.GetInstruction(window, canvases, instructFile.Instructions[7], &waiting)
				waitKeyPress()
			}
			content.RemoveAll()
			content.Add(containers.Stimuli)
			content.Refresh()
			event.CaptureEscape(window)
		}

		// I9
		content.RemoveAll()
		content.Add(containers.Stimuli)
		content.Refresh()
		event.CaptureEscape(window)
		for i := 1; i <= 5; i++ {
			obj.GetPreTrainStimulus(canvases, pictureFile.OPSlice[i])
			time.Sleep(time.Second * time.Duration(showingTime))
			canvases.Picture.File = pictureFile.Mask
			canvases.Picture.Refresh()
			time.Sleep(time.Second)
			content.RemoveAll()
			content.Add(containers.Instruction)
			content.Refresh()
			correct := false
			obj.PreTrainInstruction(window, canvases, instructFile.Instructions[9], &correct, &waiting)
			waitKeyPress()
			preTrialCorrection = append(preTrialCorrection, !correct)
			if !correct {
				// I7
				obj.GetInstruction(window, canvases, instructFile.Instructions[6], &waiting)
				waitKeyPress()
			} else {
				// I8
				obj.GetInstruction(window, canvases, instructFile.Instructions[7], &waiting)
				waitKeyPress()
			}
			content.RemoveAll()
			content.Add(containers.Stimuli)
			content.Refresh()
			event.CaptureEscape(window)
		}

		// I13
		content.RemoveAll()
		content.Add(containers.Instruction)
		content.Refresh()
		if isPassed(preTrialCorrection) {
			obj.GetInstruction(window, canvases, instructFile.Instructions[12], &waiting)
			waitKeyPress()
		} else {
			exportData()
			obj.END(window, canvases, instructFile.Instructions[15], &waiting)
			waitKeyPress()
			(*window).Close()
		}

		// Showing pics
		content.RemoveAll()
		content.Add(containers.Stimuli)
		content.Refresh()
		event.CaptureEscape(window)

		// Iterate the stimuli
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
			obj.GetInstruction(window, canvases, instructFile.Instructions[13], &waiting)
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

		// Iterate the stimuli
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
		waitKeyPress()
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

// Before the return or space is pressed, the procedure will be paused
func waitKeyPress() {
	for waiting {
		continue
	}
	waiting = true
}

// Exports a csv and a zip files
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

func isPassed(preTrain []bool) bool {
	acc := 0.0
	for _, good := range preTrain {
		if good {
			acc++
		}
	}
	criterion, _ := strconv.ParseFloat(os.Getenv("PRE_TRIAL_CRITERION"), 64)
	if precision := acc / float64(len(preTrain)); precision >= criterion {
		return true
	}
	return false
}
