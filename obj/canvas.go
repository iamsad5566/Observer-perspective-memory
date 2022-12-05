package obj

import (
	"observerPerspective/material"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const width, height float32 = 1920, 1440

var stimuliMaterial material.Stimuli = material.Stimuli{}

type Canvases struct {
	Instruction *canvas.Image
	Mask        *canvas.Image
	Picture     *canvas.Image
}

func (c *Canvases) Load() {
	stimuliMaterial.Load()
	c.Instruction = canvas.NewImageFromFile(stimuliMaterial.CurrentInstruction)
	c.Instruction.SetMinSize(fyne.Size{Width: width, Height: height})
	c.Instruction.FillMode = canvas.ImageFillContain

	c.Picture = canvas.NewImageFromFile(stimuliMaterial.CurrentPicture)
	c.Picture.SetMinSize(fyne.Size{Width: width * 0.6, Height: height * 0.6})
	c.Picture.FillMode = canvas.ImageFillContain

	c.Mask = canvas.NewImageFromFile(stimuliMaterial.Mask)
	c.Mask.SetMinSize(fyne.Size{Width: width, Height: height})
	c.Mask.FillMode = canvas.ImageFillStretch
}
