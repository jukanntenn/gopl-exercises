package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	const (
		width, height = 1024, 1024
	)
	params := map[string]float64{
		"xmin": -2,
		"xmax": 2,
		"ymin": -2,
		"ymax": 2,
		"zoom": 1,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for name := range params {
			s := r.FormValue(name)
			if s == "" {
				continue
			}
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("query param %s: %s", name, err), http.StatusBadRequest)
				return
			}
			params[name] = f
		}
		if params["xmax"] <= params["xmin"] || params["ymax"] <= params["ymin"] {
			http.Error(w, "min coordinates must be less than max", http.StatusBadRequest)
			return
		}

		xmin := params["xmin"]
		xmax := params["xmax"]
		ymin := params["ymin"]
		ymax := params["ymax"]
		zoom := params["zoom"]

		lenX := xmax - xmin
		midX := xmin + lenX/2
		xmin = midX - lenX/2/zoom
		xmax = midX + lenX/2/zoom
		lenY := ymax - ymin
		midY := ymin + lenY/2
		ymin = midY - lenY/2/zoom
		ymax = midY + lenY/2/zoom

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}
		err := png.Encode(w, img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrot(z complex128) color.Color {
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
