package loading

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// bgImageLoad represents the image that will load in the background
type bgImageLoad struct {
	uri fyne.URI
	img *canvas.Image
}

// Loads queues the images to be loaded
var Loads = make(chan bgImageLoad, 1024)
