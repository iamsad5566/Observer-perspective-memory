package event

import (
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

const percentage float32 = 1.02
const (
	FP = 5
	OP = 9
)

func CaptureEscape(window *fyne.Window) {
	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Escape" {
				(*window).Close()
			}
		})
	}
}

func CaptureZoom(window *fyne.Window, target *canvas.Image, waiting *bool, ratio float32, result *[]float32) {
	width := target.MinSize().Width
	height := target.MinSize().Height

	_, w := getHeightAndWidth()
	innerSide64, _ := strconv.ParseFloat(os.Getenv("INNER_WINDOW_SIDE"), 32)
	innerSide := float32(innerSide64)

	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Up" && width <= w/percentage {
				ratio *= percentage * percentage
				width *= percentage
				height *= percentage
				target.SetMinSize(fyne.Size{Width: width, Height: height})
				target.Refresh()
			} else if ke.Name == "Down" && height >= innerSide*percentage {
				ratio /= percentage * percentage
				width /= percentage
				height /= percentage
				target.SetMinSize(fyne.Size{Width: width, Height: height})
				target.Refresh()
			} else if ke.Name == "Return" {
				*result = append(*result, ratio)
				*waiting = false
			} else if ke.Name == "Escape" {
				(*window).Close()
			}
		})
	}
}

func SpaceContinue(window *fyne.Window, waiting *bool) {
	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Space" || ke.Name == "Return" {
				*waiting = false
			} else if ke.Name == "Escape" {
				(*window).Close()
			}
		})
	}
}

func PreTrainResponse(window *fyne.Window, correct *bool, waiting *bool, preType int) {
	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "F" && preType == FP {
				*correct = true
				*waiting = false
			} else if ke.Name == "J" && preType == OP {
				*correct = true
				*waiting = false
			} else if ke.Name == "Escape" {
				(*window).Close()
			} else if ke.Name == "F" || ke.Name == "J" {
				*correct = false
				*waiting = false
			}
		})
	}
}

func Leave(window *fyne.Window, waiting *bool) {
	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Return" || ke.Name == "Space" {
				defer (*window).Close()
			}
		})
	}
}

func getHeightAndWidth() (float32, float32) {
	h := os.Getenv("HEIGHT")
	w := os.Getenv("WIDTH")
	height, _ := strconv.ParseFloat(h, 32)
	width, _ := strconv.ParseFloat(w, 32)
	return float32(height), float32(width)
}
