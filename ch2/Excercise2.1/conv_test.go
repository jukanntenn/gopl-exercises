package tempconv

import (
	"math"
	"testing"
)

var EPSILON float64 = 0.00000001

func TestCToF(t *testing.T) {
	tests := []struct {
		input Celsius
		want  Fahrenheit
	}{
		{-100, -148},
		{0, 32},
		{100, 212},
	}

	for _, test := range tests {
		if got := CToF(test.input); got != test.want {
			t.Errorf("CToF(%q) = %v", test.input, got)
		}
	}
}

func TestFToC(t *testing.T) {
	tests := []struct {
		input Fahrenheit
		want  Celsius
	}{
		{-148, -100},
		{32, 0},
		{212, 100},
	}

	for _, test := range tests {
		if got := FToC(test.input); got != test.want {
			t.Errorf("FToC(%q) = %v", test.input, got)
		}
	}
}

func TestKToC(t *testing.T) {
	tests := []struct {
		input Kelvin
		want  Celsius
	}{
		{0, -273.15},
		{100, -173.15},
		{273.15, 0},
		{373.15, 100},
	}

	for _, test := range tests {
		if got := KToC(test.input); math.Abs(float64(got-test.want)) >= EPSILON {
			t.Errorf("KToC(%q) = %v", test.input, got)
		}
	}
}

func TestCToK(t *testing.T) {
	tests := []struct {
		input Celsius
		want  Kelvin
	}{
		{-273.15, 0},
		{-173.15, 100},
		{0, 273.15},
		{100, 373.15},
	}

	for _, test := range tests {
		if got := CToK(test.input); math.Abs(float64(got-test.want)) >= EPSILON {
			t.Errorf("CToK(%q) = %v", test.input, got)
		}
	}
}

func TestKToF(t *testing.T) {
	tests := []struct {
		input Kelvin
		want  Fahrenheit
	}{
		{0, -459.67},
		{100, -279.67},
		{999, 1338.53},
		{100000, 179540.33},
	}

	for _, test := range tests {
		if got := KToF(test.input); math.Abs(float64(got-test.want)) >= EPSILON {
			t.Errorf("KToF(%q) = %v", test.input, got)
		}
	}
}

func TestFToK(t *testing.T) {
	tests := []struct {
		input Fahrenheit
		want  Kelvin
	}{
		{-459.67, 0},
		{0, 255.3722222222222},
		{999, 810.3722222222223},
		{100000, 55810.927777777775},
	}

	for _, test := range tests {
		if got := FToK(test.input); math.Abs(float64(got-test.want)) >= EPSILON {
			t.Errorf("FToK(%q) = %v", test.input, got)
		}
	}
}
