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

// loads queues the images to be loaded
var loads = make(chan bgImageLoad, 1024)
