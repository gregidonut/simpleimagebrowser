package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"strconv"
)

func main() {
	a := app.New()

	w := a.NewWindow("simpleimagebrowser")

	w.SetContent(makeImageGrid())
	//w.Resize(fyne.NewSize(600, 600))
	//w.SetFixedSize(true)

	w.ShowAndRun()
}

// makeImageItem sets up the image item to have a label that is centered in its container
// then returns a border layout container where at the bottom the label is placed and the
// image which will consume the remaining space(which in this case will be the top since
// only the bottom of the layout is specified to be used)
func makeImageItem(imgLabel string) fyne.CanvasObject {
	label := canvas.NewText(imgLabel, color.NRGBA{R: 0xf7, G: 0xff, B: 0x5e, A: 0xff})
	label.Alignment = fyne.TextAlignCenter

	img := canvas.NewRectangle(color.NRGBA{R: 0x8b, G: 0xff, B: 0xfe, A: 0xff})
	return container.NewBorder(nil, label, nil, nil, img)
}

// makeImageGrid creates an image grid container using the GridWrap layout with a images
// laid out in it to simulate a thumbnail style presentation of a list of images calling
// the makeImageItem to present the images with a fixed cell size
func makeImageGrid() fyne.CanvasObject {
	items := make([]fyne.CanvasObject, 0)

	// loop through images to show as thumbnails
	for i := 1; i <= 10; i++ {
		mockImageName := strconv.Itoa(i)
		items = append(items, makeImageItem(mockImageName))
	}

	cellSize := fyne.NewSize(160, 160)
	return container.NewGridWrap(cellSize, items...)
}
