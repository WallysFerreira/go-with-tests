package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("imcrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		expected := 3

		counter.Inc()
		counter.Inc()
		counter.Inc()

		got := counter.Value()

		if got != expected {
			t.Errorf("Got %d, expected %d", got, expected)
		}
	})
}
