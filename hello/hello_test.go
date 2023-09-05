package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("world", "")
		want := "Hello, world"

		assertCorrectMessages(t, got, want)
	})

	t.Run("say 'Hello, World' when empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessages(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Teste", "Spanish")
		want := "Hola, Teste"

		assertCorrectMessages(t, got, want)
	})

	t.Run("in French", func (t *testing.T) {
		got := Hello("Teste", "French")
		want := "Bonjour, Teste"

		assertCorrectMessages(t, got, want)
	})
}

func assertCorrectMessages(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
