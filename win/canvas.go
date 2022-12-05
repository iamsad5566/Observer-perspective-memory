package win

import (
	"fmt"
	"observerPerspective/material"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const width, height float32 = 1920, 1440

var stimuliMaterial material.Stimuli = material.Stimuli{}

func BuildInstructWin() *fyne.Container {
	stimuliMaterial.Load()
	s := canvas.NewImageFromFile(stimuliMaterial.Description)
	s.SetMinSize(fyne.Size{Width: width, Height: height})
	s.FillMode = canvas.ImageFillContain

	return container.NewCenter(s)
}

func BuildStimuliWin(currIndex int, condition int) fyne.CanvasObject {
	stimuliMaterial.Load()
	currStimulus := Stimulus{}
	currStimulus.initialize(stimuliMaterial.Array[0])
	target := canvas.NewImageFromFile(currStimulus.Current)
	target.SetMinSize(fyne.Size{Width: width, Height: height})
	target.FillMode = canvas.ImageFillContain

	mask := canvas.NewImageFromFile(stimuliMaterial.Mask)
	mask.SetMinSize(fyne.Size{Width: width, Height: height})
	mask.FillMode = canvas.ImageFillStretch

	// insert condition == 1
	if condition == 1 {
		fmt.Println("ok")
	}

	return container.NewCenter(target, mask)
}
