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
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Area(rectangle)
		expected := 100.0

		if got != expected {
			t.Errorf("Got %.2f expected %.2f", got, expected)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		expected := 314.1592653589793

		if got != expected {
			t.Errorf("Got %.2f expected %.2f", got, expected)
		}
	})
}
