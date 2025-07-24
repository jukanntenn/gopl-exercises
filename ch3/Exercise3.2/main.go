package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange...+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type zFunc = func(float64, float64) float64

func draw(fname string, fn zFunc) {
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, fn)
			bx, by := corner(i, j, fn)
			cx, cy := corner(i, j+1, fn)
			dx, dy := corner(i+1, j+1, fn)
			if anyinf(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			fmt.Fprintf(f, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(f, "</svg>")
}

func corner(i, j int, fn zFunc) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := fn(x, y)
	if math.IsInf(z, 0) {
		return math.Inf(0), math.Inf(0)
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func moguls(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 30
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

func anyinf(args ...float64) bool {
	for _, arg := range args {
		if math.IsInf(arg, 0) {
			return true
		}
	}
	return false
}

func main() {
	draw("eggbox.svg", eggbox)
	draw("moguls.svg", moguls)
	draw("saddle.svg", saddle)
}
