package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, got, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("a")
		expected := "Could not find the word you were looking for"

		if err == nil {
			t.Fatal("Expected an error but didn't get one")
		}

		assertStrings(t, err.Error(), expected)
	})
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("Got %s expected %s", got, expected)
	}
}
