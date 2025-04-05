package main

import (
	"fmt"
	"image/png"
	"os"

	ascii "github.com/qeesung/image2ascii/convert"
)

func main() {
	f, err := os.Open("images/puppy1.png")
	if err != nil {
		panic(err)
	}
	image, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	converter := ascii.NewImageConverter()
	fmt.Printf("%s\n", converter.Image2ASCIIString(image, &ascii.DefaultOptions))
}
