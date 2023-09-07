package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name     string
		shape    ShapeWithPerimeter
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, expected: 40.0},
		{name: "Circle", shape: Circle{10}, expected: 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()

			if got != tt.expected {
				t.Errorf("%#v got %g expected %g", tt.shape, got, tt.expected)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, expected: 100.0},
		{name: "Circle", shape: Circle{10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.expected {
				t.Errorf("%#v got %g expected %g", tt.shape, got, tt.expected)
			}
		})

	}
}
