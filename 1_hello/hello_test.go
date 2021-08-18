package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB /* an interface that *testing.T & *testing.B both satisfy*/, got, want string) {
		t.Helper() // tell the test suit that this method is a helper
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Guntoro", "")

		want := "Hello, Guntoro"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")

		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")

		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Yudhy", "French")

		want := "Bonjour, Yudhy"

		assertCorrectMessage(t, got, want)
	})
}
