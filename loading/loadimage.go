package loading

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"image/color"
	"log"
)

// LoadImage is ultimately a wrapper for the canvas.NewImageFromResource that is used to
// display the image from the URI in it's container
func LoadImage(u fyne.URI) fyne.CanvasObject {
	read, err := storage.Reader(u)
	if err != nil {
		log.Println("Error opening image", err)
		return canvas.NewRectangle(color.NRGBA{R: 0xff, G: 0xaf, B: 0x62, A: 0xff})
	}

	res, err := storage.LoadResourceFromURI(read.URI())
	if err != nil {
		log.Println("Error reading image", err)
		return canvas.NewRectangle(color.NRGBA{R: 0xff, G: 0xaf, B: 0x62, A: 0xff})
	}

	img := canvas.NewImageFromResource(res)
	img.FillMode = canvas.ImageFillContain
	return img
}
