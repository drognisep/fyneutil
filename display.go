package util

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func BackgroundColor(c color.Color, obj fyne.CanvasObject) fyne.CanvasObject {
	rect := canvas.NewRectangle(c)
	return container.NewMax(rect, obj)
}
