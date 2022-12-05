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

func GetInstruction(window *fyne.Window, canvases *Canvases) {
	event.Test(window, canvases.Picture)
}
