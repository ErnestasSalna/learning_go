package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ernestas")
	want := "Hello, Ernestas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
