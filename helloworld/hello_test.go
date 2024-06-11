package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", "english")
		want := "hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := hello("Chris", "spanish")
		want := "hola Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello world if name is not provided", func(t *testing.T) {
		got := hello("", "english")
		want := "hello world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
