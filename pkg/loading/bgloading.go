package loading

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image"
)

// BgImageLoad represents the image that will load in the background
type BgImageLoad struct {
	uri fyne.URI
	img *canvas.Image
}

// LoadedImage represents the components of the image item from the layout
// to actually present the image thumbnail. these two items need be fused
// together to actually load the item which will be the responsibility of the
// for loop reading from the results channel of the corresponding worker pool
type LoadedImage struct {
	decoded image.Image
	img     *canvas.Image
}
