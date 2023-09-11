package main

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	expected := fastURL
	got := Racer(slowURL, fastURL)

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
