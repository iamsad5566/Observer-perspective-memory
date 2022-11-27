package event

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func CaptureEscape(window fyne.Window) {
	if desk, ok := window.Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(func(ke *fyne.KeyEvent) {
			if ke.Name == "Escape" {
				window.Close()
			}
		})
	}
}
