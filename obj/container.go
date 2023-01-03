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

func GetPreTrainStimulus(window *fyne.Window, canvases *Canvases, fileName string) {
	canvases.Instruction.File = fileName
	canvases.Instruction.Refresh()
}

func PreTrainInstruction(window *fyne.Window, canvases *Canvases, fileName string, correct *bool, waiting *bool) {
	canvases.Instruction.File = fileName
	canvases.Instruction.Refresh()

	event.PreTrainResponse(window, correct, waiting)
}

func GetStimulus(canvases *Canvases, fileName string) {
	canvases.Picture.File = fileName
	canvases.Picture.Refresh()
}

func GetResponseToStimulus(window *fyne.Window, canvases *Canvases, fileName string, waiting *bool, result *[]float32, ratioIndex int) {
	canvases.Picture.File = fileName
	canvases.ReSize(ratioIndex)
	canvases.Picture.Refresh()
	event.CaptureZoom(window, canvases.Picture, waiting, result)
}

func END(window *fyne.Window, canvases *Canvases, fileName string, waiting *bool) {
	canvases.Instruction.File = fileName
	canvases.Instruction.Refresh()

	event.Leave(window, waiting)
}
