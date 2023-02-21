package main

import (
	"flag"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	a := app.New()

	w := a.NewWindow("simpleimagebrowser")

	w.SetContent(makeUI())
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

// makeStatus will show more information about the focused image from the makeImageGrid
// container
func makeStatus() fyne.CanvasObject {
	return canvas.NewText("status", color.NRGBA{R: 0xc7, G: 0x5a, B: 0xff, A: 0xff})
}

// makeUI will be the main container that will contain everything
func makeUI() fyne.CanvasObject {
	status := makeStatus()
	content := makeImageGrid()
	return container.NewBorder(nil, status, nil, nil, content)
}

// startDirectory will parse commandline args and see if it can use
// the argument as a fyne.ListableURI or use the current working directory
// if not
func startDirectory() fyne.ListableURI {
	flag.Parse()
	if len(flag.Args()) < 1 {
		cwd, _ := os.Getwd()
		list, _ := storage.ListerForURI(storage.NewFileURI(cwd))
		fmt.Printf("%#v\n", list)
		return list
	}

	dir, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		cwd, _ := os.Getwd()
		log.Printf("Could not find directory: %q\nOpening: %q instead", dir, cwd)
		list, _ := storage.ListerForURI(storage.NewFileURI(cwd))
		fmt.Printf("%#v\n", list)
		return list
	}

	list, _ := storage.ListerForURI(storage.NewFileURI(dir))
	return list
}
