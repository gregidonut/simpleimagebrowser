package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")

	w.SetContent(canvas.NewText("This works!", color.White))

	w.ShowAndRun()
}
