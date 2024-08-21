package main

import "testing"

func TestHello(t *testing.T) {
	got := printHello("Brian")
	want := "Hello, Brian"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
