package main

func Perimeter(height, width float64) float64 {
	return 2 * (height + width)
}

func Area(height, width float64) float64 {
	return height * width
}

type Rectangle struct {
	Width  float64
	Height float64
}
