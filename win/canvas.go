package win

import (
	"math/rand"
	"observerPerspective/event"
	"observerPerspective/material"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const width, height float32 = 1920, 1440

var stimuliMaterial material.Stimuli = material.Stimuli{}

func BuildInstructWin(window fyne.Window) *fyne.Container {
	stimuliMaterial.Load()
	s := canvas.NewImageFromFile(stimuliMaterial.Description)
	s.SetMinSize(fyne.Size{Width: width, Height: height})
	s.FillMode = canvas.ImageFillContain

	return container.NewCenter(s)
}

func BuildShowingPicsWin(window *fyne.Window, ShowingPhase int) (fyne.CanvasObject, []string) {
	stimuliList := make([]string, 20)
	stimuliMaterial.Load()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(stimuliMaterial.Array), func(i, j int) {
		stimuliMaterial.Array[i], stimuliMaterial.Array[j] = stimuliMaterial.Array[j], stimuliMaterial.Array[i]
	})

	for i := 0; i < 20; i++ {
		stimuliList[i] = stimuliMaterial.Array[i]
	}

	currStimulus := Stimulus{}
	currStimulus.initialize(stimuliList[0])

	target := canvas.NewImageFromFile(currStimulus.Current)
	target.SetMinSize(fyne.Size{Width: width, Height: height})
	target.FillMode = canvas.ImageFillContain

	mask := canvas.NewImageFromFile(stimuliMaterial.Mask)
	mask.SetMinSize(fyne.Size{Width: width, Height: height})
	mask.FillMode = canvas.ImageFillStretch

	event.CaptureZoom(window, ShowingPhase, target)

	return container.NewCenter(target, mask), stimuliList
}
