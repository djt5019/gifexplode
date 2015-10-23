package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"log"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Missing the required filename parameter\n")
	}

	filename := os.Args[1]
	var wg sync.WaitGroup

	fh, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open '%s' - {Error: %v}\n", filename, err)
	}

	gifs, err := gif.DecodeAll(fh)
	if err != nil {
		log.Fatalf("Failed to decode '%s' - {Error: %v}\n", filename, err)
	}

	wg.Add(len(gifs.Image))

	for idx, image := range gifs.Image {
		newFileName := fmt.Sprintf("%s.%d.png", filename, idx)
		go writeImage(&wg, image, newFileName)
	}

	wg.Wait()
}

func writeImage(wg *sync.WaitGroup, image image.Image, filename string) {
	defer wg.Done()
	newFh, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create the exploded GIF frame - {Error: %v}\n", err)
	}

	err = png.Encode(newFh, image)
	if err != nil {
		log.Fatalf("Failed save the exploded GIF frame - {Error: %v}\n", err)
	}
}
