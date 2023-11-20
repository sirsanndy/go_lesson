package main

import "testing"

func TestPointer(t *testing.T) {
	var text string = "Sandy"
	t.Run("saying Mr. "+text, func(t *testing.T) {
		sandy := Man{text}
		sandy.Married()
		got := sandy.Name
		want := "Mr. " + text

		assertCorrectString(t, got, want)
	})
}

func assertCorrectString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
