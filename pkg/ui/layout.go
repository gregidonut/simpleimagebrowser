package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/gregidonut/simpleimagebrowser/pkg/listing"
	"github.com/gregidonut/simpleimagebrowser/pkg/loading"
	"image/color"
	"log"
	"os"
	"runtime"
)

// makeImageItem sets up the image item to have a label that is centered in its container
// then returns a border ui container where at the bottom the label is placed and the
// image which will consume the remaining space(which in this case will be the top since
// only the bottom of the ui is specified to be used)
func makeImageItem(u fyne.URI, jobs chan<- loading.BgImageLoad) fyne.CanvasObject {
	label := canvas.NewText(u.Name(), color.NRGBA{R: 0xf7, G: 0xff, B: 0x5e, A: 0xff})
	label.Alignment = fyne.TextAlignCenter

	img := loading.LoadImage(u, jobs)
	return container.NewBorder(nil, label, nil, nil, img)
}

// makeImageGrid creates an image grid container using the GridWrap ui with a images
// laid out in it to simulate a thumbnail style presentation of a list of images calling
// the makeImageItem to present the images with a fixed cell size
func makeImageGrid(images []fyne.URI) fyne.CanvasObject {
	items := make([]fyne.CanvasObject, 0)

	// {{ worker-pools pattern
	var workers int // sanity checking to make sure workers is an int(1000% not necessary)
	workers = runtime.NumCPU() / 2

	results := make(chan loading.LoadedImage, workers)
	jobs := make(chan loading.BgImageLoad, len(images))
	defer close(jobs)

	for i := 0; i < workers; i++ {
		go loading.DoLoadImages(jobs, results)
		go loading.RefreshImages(results)
	}

	// loop through blank images to show as thumb thumbnails while waiting for background
	// loading of images
	for _, u := range images {
		items = append(items, makeImageItem(u, jobs))
	}
	//}}

	cellSize := fyne.NewSize(160, 160)

	// wrap teh container in a scrolling parent (as an expected feature in any browser)
	imageGrid := container.NewGridWrap(cellSize, items...)
	return container.NewVScroll(imageGrid)
}

// makeStatus will show more information about the dir from the makeImageGrid container is presenting
func makeStatus(images []fyne.URI) fyne.CanvasObject {
	dirPath, _ := os.Getwd()
	status := fmt.Sprintf("Directory: %s; %d items", dirPath, len(images))
	return canvas.NewText(status, color.NRGBA{R: 0xc7, G: 0x5a, B: 0xff, A: 0xff})
}

// MakeUI will be the main container that will contain everything
func MakeUI(dir fyne.ListableURI) fyne.CanvasObject {
	list, err := dir.List()
	if err != nil {
		log.Println("Error listing directory", err)
	}

	images := listing.FilterImages(list)

	status := makeStatus(images)
	content := makeImageGrid(images)
	return container.NewBorder(nil, status, nil, nil, content)
}
