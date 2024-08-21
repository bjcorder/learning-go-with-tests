package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := printHello("Brian", "")
		want := "Hello, Brian"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := printHello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := printHello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := printHello("Jacques", "French")
		want := "Bonjour, Jacques"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
