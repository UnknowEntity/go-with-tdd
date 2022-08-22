package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello to People", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello World when omit name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello to People in Spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello to People in French", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour Chris"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
