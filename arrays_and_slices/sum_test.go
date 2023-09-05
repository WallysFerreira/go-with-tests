package main

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	expected := 6

	if got != expected {
		t.Errorf("Got %d expected %d, %v", got, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if got != expected {
		t.Errorf("Got %v expected %v", got, expected)
	}
}
