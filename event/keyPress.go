package event

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

const percentage float32 = 1.03

func CaptureEscape(window *fyne.Window) {
	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Escape" {
				(*window).Close()
			}
		})
	}
}

func CaptureZoom(window *fyne.Window, target *canvas.Image, waiting *bool) {
	width := target.MinSize().Width
	height := target.MinSize().Height

	if desk, ok := (*window).Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Up" && height <= 1440/percentage {
				width *= percentage
				height *= percentage
				target.SetMinSize(fyne.Size{Width: width, Height: height})
				target.Refresh()
			} else if ke.Name == "Down" && height >= 640*percentage {
				width /= percentage
				height /= percentage
				target.SetMinSize(fyne.Size{Width: width, Height: height})
				target.Refresh()
			} else if ke.Name == "Return" {
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
