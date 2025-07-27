package main

import (
	"image/color"
	"testing"
)

func benchmarkMandelbrot(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

// BenchmarkMandelbrotComplex64-32         128925187                9.331 ns/op
func BenchmarkMandelbrotComplex64(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrot64)
}

// BenchmarkMandelbrotComplex128-32        137808424                8.693 ns/op
func BenchmarkMandelbrotComplex128(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrot128)
}

// BenchmarkMandelbrotBigFloat-32           7979269               150.6 ns/op
func BenchmarkMandelbrotBigFloat(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrotBigFloat)
}

// BenchmarkMandelbrotBigRat-32             1995380               590.1 ns/op
func BenchmarkMandelbrotBigRat(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrotBigRat)
}
