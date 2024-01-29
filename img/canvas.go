package canvas

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/githubmo/go-mandelbrot/compute"
)

const (
	MinLength = 100
	MaxLength = 1000
)

func GenerateImage(width int, height int, min complex128, max complex128) (*image.RGBA, error) {
	if width < MinLength || width > MaxLength || height < MinLength || height > MaxLength {
		errorMsg := fmt.Sprintf("%d width and %d height are not supported, supported sizes are 100x100 up to 1000x1000",
			width, height)
		err := errors.New(errorMsg)
		fmt.Println(err)
		return nil, err
	}

	mandelbrotImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	realStep := (real(max) - real(min)) / float64(width)
	imaginaryStep := (imag(max) - imag(min)) / float64(height)

	// go 1.22 will support this
	// for i := range width {
	// 	for j := range height {
	// 		fmt.Printf("%d %d\n", width, height)
	// 	}
	// }

	for i := 0; i <= width; i++ {
		for j := 0; j <= height; j++ {
			cr := real(min) + realStep*float64(i)
			ci := imag(min) + imaginaryStep*float64(j)

			result := compute.Compute(complex128(complex(cr, ci)))
			shade := color.RGBA{result, result, result, 0xff}
			mandelbrotImage.SetRGBA(i, j, shade)
		}
	}

	return mandelbrotImage, nil
}

func SaveImage(image *image.RGBA) {
	f, err := os.Create("result.png")
	if err != nil {
		fmt.Printf("Could not generate image: %v\n", err)
	}
	err = png.Encode(f, image)
	if err != nil {
		return
	}
}
