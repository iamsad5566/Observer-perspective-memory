package event

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

const percentage float32 = 1.05

func CaptureEscape(window *fyne.Window) {
	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Escape" {
				(*window).Close()
			}
		})
	}
}

func CaptureZoom(window *fyne.Window, condition int, target *canvas.Image) {
	if condition == 0 {
		fmt.Println(condition)
		return
	}

	width := target.MinSize().Width
	height := target.MinSize().Height

	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Up" {
				width *= percentage
				height *= percentage
				target.SetMinSize(fyne.Size{Width: width, Height: height})
				target.Refresh()
			} else if ke.Name == "Down" {
				width /= percentage
				height /= percentage
				target.SetMinSize(fyne.Size{Width: width, Height: height})
				target.Refresh()
			} else if ke.Name == "Escape" {
				(*window).Close()
			} else if ke.Name == "Space" {

			}
		})
	}
}

func Test(window *fyne.Window, target *canvas.Image) {
	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Space" {
				target.File = "material/P2.jpg"
				target.Refresh()
			}
		})
	}
}
