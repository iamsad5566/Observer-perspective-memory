package obj

import (
	"math/rand"
	"observerPerspective/material"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var width float32
var height float32
var picWidth float32
var picHeight float32
var ratio = []float32{1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9, 0.9}

type Canvases struct {
	Instruction *canvas.Image
	Mask        *canvas.Image
	Picture     *canvas.Image
}

func (c *Canvases) Load(instructFile *material.InstructFile, pictureFile *material.PictureFile) {
	loadEnv()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ratio), func(i, j int) { ratio[i], ratio[j] = ratio[j], ratio[i] })

	c.Instruction = canvas.NewImageFromFile(instructFile.Instructions[0])
	c.Instruction.SetMinSize(fyne.Size{Width: width, Height: height})
	c.Instruction.FillMode = canvas.ImageFillContain

	c.Picture = canvas.NewImageFromFile(pictureFile.CurrentPicture)
	c.Picture.SetMinSize(fyne.Size{Width: picWidth, Height: picHeight})
	c.Picture.FillMode = canvas.ImageFillContain

	c.Mask = canvas.NewImageFromFile(pictureFile.Mask)
	c.Mask.SetMinSize(fyne.Size{Width: width, Height: height})
	c.Mask.FillMode = canvas.ImageFillStretch
}

func (c *Canvases) ReSize(index int) float32 {
	r := ratio[index%len(ratio)]
	c.Picture.SetMinSize(fyne.Size{Width: picWidth * r , Height: picHeight * r })
	return r
}

func loadEnv() {
	strWidth := os.Getenv("WIDTH")
	strHeight := os.Getenv("HEIGHT")
	width64, _ := strconv.ParseFloat(strWidth, 32)
	height64, _ := strconv.ParseFloat(strHeight, 32)
	width = float32(width64)
	height = float32(height64)

	picW, _ := strconv.ParseFloat(os.Getenv("PIC_WIDTH"), 32)
	picH, _ := strconv.ParseFloat(os.Getenv("PIC_HEIGHT"), 32)
	picWidth = float32(picW)
	picHeight = float32(picH)
}
