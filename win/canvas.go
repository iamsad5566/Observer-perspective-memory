package win

import (
	"fmt"
	"observerPerspective/material"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func BuildInstructWin(instruct string, keypress string) fyne.CanvasObject {
	upperText := canvas.NewText(instruct, nil)
	lowerText := canvas.NewText(keypress, nil)
	upperText.TextSize = 35
	lowerText.TextSize = 30
	content := container.NewGridWithRows(4, canvas.NewLine(nil), container.NewCenter(upperText), container.NewCenter(lowerText), canvas.NewLine(nil))
	return content
}

func BuildStimuliWin(currIndex int, condition int) fyne.CanvasObject {
	var width, height float32 = 1920, 1440

	stimuli := material.Stimuli{}
	stimuli.Load()
	target := canvas.NewImageFromFile(stimuli.Array[currIndex])
	target.SetMinSize(fyne.Size{Width: width, Height: height})
	target.FillMode = canvas.ImageFillContain

	mask := canvas.NewImageFromFile(stimuli.Mask)
	mask.SetMinSize(fyne.Size{Width: width, Height: height})
	mask.FillMode = canvas.ImageFillStretch

	// insert condition == 1
	if condition == 1 {
		fmt.Println("ok")
	}

	return container.NewCenter(target, mask)
}
