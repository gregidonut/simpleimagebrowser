package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/gregidonut/simpleimagebrowser/pkg/listing"
	"github.com/gregidonut/simpleimagebrowser/pkg/loading"
	"github.com/gregidonut/simpleimagebrowser/pkg/ui"
	"runtime"
)

func main() {
	a := app.New()

	w := a.NewWindow("simpleimagebrowser")

	w.SetContent(ui.MakeUI(listing.StartDirectory()))
	//w.Resize(fyne.NewSize(600, 600))
	//w.SetFixedSize(true)

	var workers int // sanity checking to make sure workers is an int(1000% not necessary)
	workers = runtime.NumCPU() / 2
	for i := 0; i < workers; i++ {
		go loading.DoLoadImages()
	}
	w.ShowAndRun()
}
