package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

const DestinationImageFileName = "empty.png"

func saveImage(filename string, img image.Image) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outfile.Close()

	png.Encode(outfile, img)
	return nil
}

func fillPixels(pixels []uint8)

func main() {
	destinationImage := image.NewRGBA(image.Rect(0, 0, 256, 256))

	fillPixels(destinationImage.Pix)

	err := saveImage(DestinationImageFileName, destinationImage)
	if err != nil {
		log.Fatal(err)
	}
}
