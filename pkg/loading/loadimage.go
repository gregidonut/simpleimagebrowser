package loading

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"github.com/nfnt/resize"
	"image"
	"log"
)

// LoadImage will add the img and to the background channel queue
func LoadImage(u fyne.URI, jobs chan<- BgImageLoad) fyne.CanvasObject {
	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain

	// adding the image to the background channel queue
	go func() {
		jobs <- BgImageLoad{u, img}
	}()
	return img
}

// doLoadImage is responsible for reading the image file from the fyne.URI and returning
// the in-memory image to the calling container/CanvasObject
func doLoadImage(u fyne.URI, img *canvas.Image) LoadedImage {
	read, err := storage.Reader(u)
	if err != nil {
		log.Println("Error opening image", err)
		return LoadedImage{}
	}
	defer read.Close()

	raw, _, err := image.Decode(read)
	if err != nil {
		log.Println("Error decoding image", err)
		return LoadedImage{}
	}

	//img.Image = scaleImage(raw)
	//img.Refresh()
	return LoadedImage{scaleImage(raw), img}

}

// scaleImage is essentially a wrapper for the resize.Thumbnail function but specifically
// uses the resize.Lanczos3 for a more efficient scaling routine
func scaleImage(img image.Image) image.Image {
	// specifying fixed size twice the size of the actual parent container for a more
	// pixel-dense output (should probably adjust this to become smarter in the future)
	return resize.Thumbnail(320, 320, img, resize.Lanczos3)
}

// DoLoadImages is the function that will be called in main as a goroutine, it will
// range through the Loads channel and call the main logic function: doLoadImage on each load
// which will redraw the image everytime the parent container is resized since doLoadImage calls
// the Refresh() method from the *canvas.Image struct
func DoLoadImages(jobs <-chan BgImageLoad, results chan<- LoadedImage) {
	for job := range jobs {
		results <- doLoadImage(job.uri, job.img)
	}
}

func RefreshImages(results <-chan LoadedImage) {
	for res := range results {
		res.img.Image = res.decoded
		res.img.Refresh()
	}
}
