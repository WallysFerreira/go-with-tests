package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collections of any sizes", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("Got %d expected %d, %v", got, expected, numbers)
		}
	})
}
