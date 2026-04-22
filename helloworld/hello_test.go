package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Fawaz", "")
		want := "Hello, Fawaz"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Fawaz", "Spanish")
		want := "Hola, Fawaz"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Fawaz", "French")
		want := "Bonjour, Fawaz"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Yoruba", func(t *testing.T) {
		got := Hello("Fawaz", "Yoruba")
		want := "Kaaro, Fawaz"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
