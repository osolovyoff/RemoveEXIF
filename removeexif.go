//removeexif.go
// Alex Meys

package main

import (
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	photoDirectoryName := "from"
	if len(os.Args) > 1 {
		photoDirectoryName = os.Args[1]
	}

	folder, err := ioutil.ReadDir(photoDirectoryName)
	if err != nil {
		log.Fatalln("Failed", folder)
	}

	for _, file := range folder {
		pathToFolder, err := filepath.Abs(photoDirectoryName)
		if err != nil {
			log.Fatalln("Failed", err)
		}
		pathToFile := path.Join(pathToFolder, file.Name())

		fileCurrent, err := os.Open(pathToFile)
		if err != nil {
			log.Fatalln("Failed", err)
		}

		defer fileCurrent.Close()

		img, _, err := image.Decode(fileCurrent)
		if err != nil {
			log.Fatalln("Failed", err)
		}

		outfile, err := os.Create(file.Name())
		if err != nil {
			log.Fatalln("Failed", err)
		}

		defer outfile.Close()

		if err := jpeg.Encode(outfile, img, nil); err != nil {
			log.Fatalln("failed", err)
		}
	}
}
