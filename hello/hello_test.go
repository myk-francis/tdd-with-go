package main

import "testing"

func TestHello(t *testing.T)  {
	got := Hello("Chris")
	want := "Hello, Chris!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloTwo(t *testing.T)  {
	assertCorrectMessage := func (t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func (t *testing.T)  {
		got := Hello("Chris")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World!' when an empty string is provided", func (t *testing.T)  {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
}