package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello("neimv", "")
	want := "Hello neimv"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("neimv", "")
		want := "Hello neimv"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("sideris", "Spanish")
		want := "Hola sideris"
		assertCorrectMessage(t, got, want)
	})
}
