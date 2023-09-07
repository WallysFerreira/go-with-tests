package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("Got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()

		got := shape.Area()

		if got != expected {
			t.Errorf("Got %.2f expected %.2f", got, expected)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}

		checkArea(t, rectangle, 100.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		checkArea(t, circle, expected)
	})
}
