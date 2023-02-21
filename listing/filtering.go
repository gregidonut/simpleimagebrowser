package listing

import (
	"flag"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// StartDirectory will parse commandline args and see if it can use
// the argument as a fyne.ListableURI or use the current working directory
// if not
func StartDirectory() fyne.ListableURI {
	flag.Parse()
	if len(flag.Args()) < 1 {
		cwd, _ := os.Getwd()
		list, _ := storage.ListerForURI(storage.NewFileURI(cwd))
		fmt.Printf("%#v\n", list)
		return list
	}

	dir, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		cwd, _ := os.Getwd()
		log.Printf("Could not find directory: %q\nOpening: %q instead", dir, cwd)
		list, _ := storage.ListerForURI(storage.NewFileURI(cwd))
		fmt.Printf("%#v\n", list)
		return list
	}

	list, _ := storage.ListerForURI(storage.NewFileURI(dir))
	return list
}

// isImage parses the fyne.URI and returns true if the extension is in the list of supported file extensions
func isImage(file fyne.URI) bool {
	ext := strings.ToLower(file.Extension())

	listOfImgExt := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
	}

	return listOfImgExt[ext]
}

// FilterImages uses isImage to generate a list of supported extensions
func FilterImages(files []fyne.URI) []fyne.URI {
	images := make([]fyne.URI, 0)

	for _, file := range files {
		if isImage(file) {
			images = append(images, file)
		}
	}

	return images
}
