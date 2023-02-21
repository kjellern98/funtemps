package conv

import (
	"fmt"
	"math"
	"testing"
)

func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{212, 100},
		{32, 0},
		{-40, -40},
		{98.6, 37},
		{68, 20},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-4) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}

}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 100, want: 212},
		{input: 0, want: 32},
		{input: -40, want: -40},
		{input: 37, want: 98.6},
		{input: 20, want: 68},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected:  %.2f, got: %.2f",
				tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 100, want: 373.15},
		{input: 0, want: 273.15},
		{input: -40, want: 233.15},
		{input: 37, want: 310.15},
		{input: 20, want: 293.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected:  %.2f, got: %.2f",
				tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 373.15, want: 100},
		{input: 273.15, want: 0},
		{input: 233.15, want: -40},
		{input: 310.15, want: 37},
		{input: 293.15, want: 20},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected:  %.2f, got: %.2f",
				tc.want, got)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 212, want: 373.15},
		{input: 32, want: 273.15},
		{input: -40, want: 233.15},
		{input: 98.6, want: 310.15},
		{input: 68, want: 293.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected:  %.2f, got: %.2f",
				tc.want, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 373.15, want: 212},
		{input: 273.15, want: 32},
		{input: 233.15, want: -40},
		{input: 310.15, want: 98.6},
		{input: 293.15, want: 68},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected:  %.2f, got: %.2f",
				tc.want, got)
		}
	}
}

func withinTolerance(a, b, error float64) bool {
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	fmt.Println(difference)
	fmt.Println(b)

	if b == 0 {
		return difference < error
	}

	return (difference / math.Abs(b)) < error
}
