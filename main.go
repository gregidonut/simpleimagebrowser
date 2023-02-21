package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {
	a := app.New()

	w := a.NewWindow("simpleimagebrowser")

	w.SetContent(makeImageItem())
	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(true)

	w.ShowAndRun()
}

// makeImageItem sets up the image item to have a label that is centered in its container
// then returns a border layout container where at the bottom the label is placed and the
// image which will consume the remaining space(which in this case will be the top since
// only the bottom of the layout is specified to be used)
func makeImageItem() fyne.CanvasObject {
	label := canvas.NewText("label", color.NRGBA{R: 0xf7, G: 0xff, B: 0x5e, A: 0xff})
	label.Alignment = fyne.TextAlignCenter

	img := canvas.NewRectangle(color.NRGBA{R: 0x8b, G: 0xff, B: 0xfe, A: 0xff})
	return container.NewBorder(nil, label, nil, nil, img)
}
