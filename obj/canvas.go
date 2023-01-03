package obj

import (
	"observerPerspective/material"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

<<<<<<< HEAD
const width, height float32 = 1600, 1200
=======
var width float32
var height float32
var ratio float32
>>>>>>> bb06bbf39696c47d682641d8eb2d4bbc77dfe447

type Canvases struct {
	Instruction *canvas.Image
	Mask        *canvas.Image
	Picture     *canvas.Image
}

func (c *Canvases) Load(instructFile *material.InstructFile, pictureFile *material.PictureFile) {
	loadEnv()
	c.Instruction = canvas.NewImageFromFile(instructFile.CurrentInstruction)
	c.Instruction.SetMinSize(fyne.Size{Width: width, Height: height})
	c.Instruction.FillMode = canvas.ImageFillContain

	c.Picture = canvas.NewImageFromFile(pictureFile.CurrentPicture)
	c.Picture.SetMinSize(fyne.Size{Width: width * ratio, Height: height * ratio})
	c.Picture.FillMode = canvas.ImageFillContain

	c.Mask = canvas.NewImageFromFile(pictureFile.Mask)
	c.Mask.SetMinSize(fyne.Size{Width: width, Height: height})
	c.Mask.FillMode = canvas.ImageFillContain
}

func (c *Canvases) ReSize() {
	c.Picture.SetMinSize(fyne.Size{Width: width * ratio, Height: height * ratio})
}

func loadEnv() {
	strWidth := os.Getenv("WIDTH")
	strHeight := os.Getenv("LENGTH")
	ratio64, _ := strconv.ParseFloat(os.Getenv("INITIAL_PICTURE_RATIO"), 32)
	width64, _ := strconv.ParseFloat(strWidth, 32)
	height64, _ := strconv.ParseFloat(strHeight, 32)
	width = float32(width64)
	height = float32(height64)
	ratio = float32(ratio64)
}
