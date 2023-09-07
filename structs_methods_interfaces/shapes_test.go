package main

import "testing"

/*
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("Got %.2f expected %.2f", got, expected)
	}
}
*/

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{shape: Rectangle{10, 10}, expected: 100.0},
		{shape: Circle{10}, expected: 314.1592653589793},
		{shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.expected {
			t.Errorf("Got %g expected %g", got, tt.expected)
		}
	}
}
