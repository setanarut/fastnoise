package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/setanarut/fastnoise"
)

func main() {

	// Create and configure noise state (either float32 or float64)
	var noise = fastnoise.New[float32]()
	noise.NoiseType(fastnoise.Cellular)

	img := image.NewGray(image.Rect(0, 0, 512, 512))

	// Gather noise data
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			v := MapRange(noise.Noise2D(x, y), -1, 1, 0, 255)
			img.SetGray(x, y, color.Gray{uint8(v)})
		}
	}
	WritePNG("cellular.png", img)
}

func MapRange(v, a, b, c, d float32) float32 {
	return (v-a)/(b-a)*(d-c) + c
}

func WritePNG(name string, im image.Image) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, im); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
