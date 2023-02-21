package loading

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// LoadImage will add the img and to the background channel queue
func LoadImage(u fyne.URI) fyne.CanvasObject {
	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain

	// adding the image to the background channel queue
	loads <- bgImageLoad{u, img}
	return img
}
