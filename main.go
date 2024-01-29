package main

import (
	"fmt"

	canvas "github.com/githubmo/go-mandelbrot/img"
)

func main() {
	image, err := canvas.GenerateImage(740, 605, -2+-1.2i, 1+1.2i)
	if err != nil {
		fmt.Println(err)
	} else {
		canvas.SaveImage(image)
	}
}
