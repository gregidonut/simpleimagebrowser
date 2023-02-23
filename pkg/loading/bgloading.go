package loading

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// BgImageLoad represents the image that will load in the background
type BgImageLoad struct {
	uri fyne.URI
	img *canvas.Image
}
