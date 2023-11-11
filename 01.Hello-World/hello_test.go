package main

import "testing"

func TestHello(t *testing.T){

	t.Run("say hello to people", func(t *testing.T){

	got := Hello("Chris", "English")
	want := "Hello, Chris"
	assertCorrectMessage(t, got, want)
})

t.Run("Default to world when no person", func(t *testing.T){

	got := Hello("", "English")
	want := "Hello, world"

	assertCorrectMessage(t, got, want)
})

t.Run("in Spanish", func(t *testing.T){
	got := Hello("Elodie", "Spanish")
	want := "Hola, Elodie"
	assertCorrectMessage(t, got, want)
})

t.Run("in French", func(t *testing.T){
	got := Hello("Elodie", "French")
	want := "Bonjour, Elodie"
	assertCorrectMessage(t, got, want)
})

t.Run("in Other", func(t *testing.T){
	got := Hello("Elodie", "Bulgarian")
	want := "Hello, Elodie"
	assertCorrectMessage(t, got, want)
})

t.Run("in English", func(t *testing.T){
	got := Hello("Elodie", "English")
	want := "Hello, Elodie"
	assertCorrectMessage(t, got, want)
})
}


func assertCorrectMessage(t testing.TB, got string, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
