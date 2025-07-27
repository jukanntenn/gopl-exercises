package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img64 := image.NewRGBA(image.Rect(0, 0, width, height))
	img128 := image.NewRGBA(image.Rect(0, 0, width, height))
	imgBigFloat := image.NewRGBA(image.Rect(0, 0, width, height))
	// imgBigRat := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img64.Set(px, py, mandelbrot64(z))
			img128.Set(px, py, mandelbrot128(z))
			imgBigFloat.Set(px, py, mandelbrotBigFloat(z))
			// took a very long time
			// imgBigRat.Set(px, py, mandelbrotBigRat(z))
		}
	}

	file64, _ := os.Create("complex64.png")
	png.Encode(file64, img64)
	file64.Close()

	file128, _ := os.Create("complex128.png")
	png.Encode(file128, img128)
	file128.Close()

	fileBigFloat, _ := os.Create("big_float.png")
	png.Encode(fileBigFloat, imgBigFloat)
	fileBigFloat.Close()

	// fileBigRat, _ := os.Create("big_rat.png")
	// png.Encode(fileBigRat, imgBigRat)
	// fileBigRat.Close()
}

func mandelbrot64(z complex128) color.Color {
	const iterations = 200

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			r := uint8((n * 8) % 255)
			g := uint8((n * 16) % 255)
			b := uint8((n * 32) % 255)
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := uint8((n * 8) % 255)
			g := uint8((n * 16) % 255)
			b := uint8((n * 32) % 255)
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200

	zReal := big.NewFloat(real(z))
	zImag := big.NewFloat(imag(z))
	vReal := big.NewFloat(0)
	vImag := big.NewFloat(0)
	tempReal := big.NewFloat(0)
	tempImag := big.NewFloat(0)
	two := big.NewFloat(2)
	four := big.NewFloat(4)

	for n := uint8(0); n < iterations; n++ {
		tempReal.Mul(vReal, vReal)
		tempImag.Mul(vImag, vImag)
		tempReal.Sub(tempReal, tempImag)
		tempReal.Add(tempReal, zReal)

		tempImag.Mul(vReal, vImag)
		tempImag.Mul(tempImag, two)
		tempImag.Add(tempImag, zImag)

		vReal.Set(tempReal)
		vImag.Set(tempImag)

		tempReal.Mul(vReal, vReal)
		tempImag.Mul(vImag, vImag)
		tempReal.Add(tempReal, tempImag)

		if tempReal.Cmp(four) > 0 {
			r := uint8((n * 8) % 255)
			g := uint8((n * 16) % 255)
			b := uint8((n * 32) % 255)
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrotBigRat(z complex128) color.Color {
	const iterations = 200

	zReal := big.NewRat(int64(real(z)*1000000), 1000000)
	zImag := big.NewRat(int64(imag(z)*1000000), 1000000)
	vReal := big.NewRat(0, 1)
	vImag := big.NewRat(0, 1)
	tempReal := big.NewRat(0, 1)
	tempImag := big.NewRat(0, 1)
	two := big.NewRat(2, 1)
	four := big.NewRat(4, 1)

	for n := uint8(0); n < iterations; n++ {
		tempReal.Mul(vReal, vReal)
		tempImag.Mul(vImag, vImag)
		tempReal.Sub(tempReal, tempImag)
		tempReal.Add(tempReal, zReal)

		tempImag.Mul(vReal, vImag)
		tempImag.Mul(tempImag, two)
		tempImag.Add(tempImag, zImag)

		vReal.Set(tempReal)
		vImag.Set(tempImag)

		tempReal.Mul(vReal, vReal)
		tempImag.Mul(vImag, vImag)
		tempReal.Add(tempReal, tempImag)

		if tempReal.Cmp(four) > 0 {
			r := uint8((n * 8) % 255)
			g := uint8((n * 16) % 255)
			b := uint8((n * 32) % 255)
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
