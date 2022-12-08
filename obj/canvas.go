package obj

import (
	"observerPerspective/material"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const width, height float32 = 1600, 1200

type Canvases struct {
	Instruction *canvas.Image
	Mask        *canvas.Image
	Picture     *canvas.Image
}

func (c *Canvases) Load(instructFile *material.InstructFile, pictureFile *material.PictureFile) {
	c.Instruction = canvas.NewImageFromFile(instructFile.CurrentInstruction)
	c.Instruction.SetMinSize(fyne.Size{Width: width, Height: height})
	c.Instruction.FillMode = canvas.ImageFillContain

	c.Picture = canvas.NewImageFromFile(pictureFile.CurrentPicture)
	c.Picture.SetMinSize(fyne.Size{Width: width * 0.7, Height: height * 0.7})
	c.Picture.FillMode = canvas.ImageFillContain

	c.Mask = canvas.NewImageFromFile(pictureFile.Mask)
	c.Mask.SetMinSize(fyne.Size{Width: width, Height: height})
	c.Mask.FillMode = canvas.ImageFillContain
}
