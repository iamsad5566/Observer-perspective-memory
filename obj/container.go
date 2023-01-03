package obj

import (
	"observerPerspective/event"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Containers struct {
	Instruction *fyne.Container
	Stimuli     *fyne.Container
}

func (c *Containers) Load(canvases *Canvases) {
	c.Instruction = container.NewCenter(canvases.Instruction)
	c.Stimuli = container.NewCenter(canvases.Picture, canvases.Mask)
}

func GetInstruction(window *fyne.Window, canvases *Canvases, fileName string, waiting *bool) {
	canvases.Instruction.File = fileName
	canvases.Instruction.Refresh()

	event.SpaceContinue(window, waiting)
}

func GetStimulus(canvases *Canvases, fileName string) {
	canvases.Picture.File = fileName
	canvases.Picture.Refresh()
}

func GetResponseToStimulus(window *fyne.Window, canvases *Canvases, fileName string, waiting *bool, result *[]float32) {
	canvases.Picture.File = fileName
	canvases.ReSize()
	canvases.Picture.Refresh()
	event.CaptureZoom(window, canvases.Picture, waiting, result)
}
