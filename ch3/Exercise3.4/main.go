package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange...+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
	blue    = "#0000ff"
)

// http://localhost:8000/?color=%23ff0000&width=800&height=800
// Note: escape hash character in URL
func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		height := 600
		if r.Form["height"] != nil {
			var err error
			height, err = strconv.Atoi(r.Form["height"][0])
			if err != nil {
				height = 600
			}
		}

		width := 320
		if r.Form["width"] != nil {
			var err error
			width, err = strconv.Atoi(r.Form["width"][0])
			if err != nil {
				height = 320
			}
		}

		color := blue
		if r.Form["color"] != nil {
			color = r.Form["color"][0]
		}
		fmt.Println(height, width, color)
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w, height, width, color)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func svg(out io.Writer, height int, width int, color string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, height, width, xyscale, zscale)
			bx, by := corner(i, j, height, width, xyscale, zscale)
			cx, cy := corner(i, j+1, height, width, xyscale, zscale)
			dx, dy := corner(i+1, j+1, height, width, xyscale, zscale)
			fmt.Fprintf(out, "<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(out, "</svg>")
}

func corner(i, j, height, width int, xyscale float64, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
