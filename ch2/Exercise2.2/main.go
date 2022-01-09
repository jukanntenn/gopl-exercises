package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	tempconv "github.com/jukanntenn/gopl-exercises/ch2/Exercise2.1"
)

func main() {
	numbers := make([]float64, 0)
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			n, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			numbers = append(numbers, n)
		}
	} else {
		for _, arg := range os.Args[1:] {
			n, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			numbers = append(numbers, n)
		}
	}

	for _, n := range numbers {
		fah := tempconv.Fahrenheit(n)
		cel := tempconv.Celsius(n)
		fmt.Printf("%s = %s, %s = %s\n", fah, tempconv.FToC(fah), cel, tempconv.CToF(cel))

		fee := Feet(n)
		met := Meter(n)
		fmt.Printf("%s = %s, %s = %s\n", fee, FToM(fee), met, MToF(met))

		pou := Pound(n)
		kil := Kilogram(n)
		fmt.Printf("%s = %s, %s = %s\n", pou, PToK(pou), kil, KToP(kil))
	}
}

type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

func FToM(f Feet) Meter {
	return Meter(f / 3.2808)
}

func MToF(m Meter) Feet {
	return Feet(m * 3.2808)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p / 2.2046)
}

func KToP(k Kilogram) Pound {
	return Pound(k * 2.2046)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gFeet", f)
}
func (m Meter) String() string {
	return fmt.Sprintf("%gMeter", m)
}
func (p Pound) String() string {
	return fmt.Sprintf("%gPound", p)
}
func (k Kilogram) String() string {
	return fmt.Sprintf("%gKilogram", k)
}
