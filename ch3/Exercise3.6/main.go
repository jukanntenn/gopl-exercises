package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var r, g, b, a uint32

			for dy := 0; dy < 2; dy++ {
				for dx := 0; dx < 2; dx++ {
					subX := float64(px) + float64(dx)*0.5 + 0.25
					subY := float64(py) + float64(dy)*0.5 + 0.25

					x := subX/width*(xmax-xmin) + xmin
					y := subY/height*(ymax-ymin) + ymin
					z := complex(x, y)

					c := mandelbrot(z)
					sr, sg, sb, sa := c.RGBA()

					r += sr
					g += sg
					b += sb
					a += sa
				}
			}

			avgColor := color.RGBA64{
				R: uint16(r / 4),
				G: uint16(g / 4),
				B: uint16(b / 4),
				A: uint16(a / 4),
			}

			img.Set(px, py, avgColor)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
