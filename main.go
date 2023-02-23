package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/gregidonut/simpleimagebrowser/pkg/listing"
	"github.com/gregidonut/simpleimagebrowser/pkg/ui"
)

func main() {
	a := app.New()

	w := a.NewWindow("simpleimagebrowser")

	w.SetContent(ui.MakeUI(listing.StartDirectory()))
	//w.Resize(fyne.NewSize(600, 600))
	//w.SetFixedSize(true)

	w.ShowAndRun()
}
