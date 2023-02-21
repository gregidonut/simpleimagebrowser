package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/gregidonut/simpleimagebrowser/listing"
	"github.com/gregidonut/simpleimagebrowser/loading"
	"image/color"
	"log"
	"os"
)

func main() {
	a := app.New()

	w := a.NewWindow("simpleimagebrowser")

	w.SetContent(makeUI(listing.StartDirectory()))
	//w.Resize(fyne.NewSize(600, 600))
	//w.SetFixedSize(true)

	go loading.DoLoadImages()
	w.ShowAndRun()
}

// makeImageItem sets up the image item to have a label that is centered in its container
// then returns a border layout container where at the bottom the label is placed and the
// image which will consume the remaining space(which in this case will be the top since
// only the bottom of the layout is specified to be used)
func makeImageItem(u fyne.URI) fyne.CanvasObject {
	label := canvas.NewText(u.Name(), color.NRGBA{R: 0xf7, G: 0xff, B: 0x5e, A: 0xff})
	label.Alignment = fyne.TextAlignCenter

	img := loading.LoadImage(u)
	return container.NewBorder(nil, label, nil, nil, img)
}

// makeImageGrid creates an image grid container using the GridWrap layout with a images
// laid out in it to simulate a thumbnail style presentation of a list of images calling
// the makeImageItem to present the images with a fixed cell size
func makeImageGrid(images []fyne.URI) fyne.CanvasObject {
	items := make([]fyne.CanvasObject, 0)

	// loop through images to show as thumbnails
	for _, u := range images {
		items = append(items, makeImageItem(u))
	}

	cellSize := fyne.NewSize(160, 160)
	return container.NewGridWrap(cellSize, items...)
}

// makeStatus will show more information about the dir from the makeImageGrid container is presenting
func makeStatus(images []fyne.URI) fyne.CanvasObject {
	dirPath, _ := os.Getwd()
	status := fmt.Sprintf("Directory: %s; %d items", dirPath, len(images))
	return canvas.NewText(status, color.NRGBA{R: 0xc7, G: 0x5a, B: 0xff, A: 0xff})
}

// makeUI will be the main container that will contain everything
func makeUI(dir fyne.ListableURI) fyne.CanvasObject {
	list, err := dir.List()
	if err != nil {
		log.Println("Error listing directory", err)
	}

	images := listing.FilterImages(list)

	status := makeStatus(images)
	content := makeImageGrid(images)
	return container.NewBorder(nil, status, nil, nil, content)
}
