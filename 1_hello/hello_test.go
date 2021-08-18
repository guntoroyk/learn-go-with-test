package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Guntoro")

	want := "Hello, Guntoro"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
