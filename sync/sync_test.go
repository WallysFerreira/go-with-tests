package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("imcrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		expected := 3

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, expected)
	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		expected := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, expected)
	})
}

func NewCounter() *Counter {
	return &Counter{}
}

func assertCounter(t testing.TB, got *Counter, expected int) {
	t.Helper()

	if got.Value() != expected {
		t.Errorf("Got %d, expected %d", got, expected)
	}
}
