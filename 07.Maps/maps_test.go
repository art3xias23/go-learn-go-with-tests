package main

import "testing"

func TestSearch(t *testing.T){

	assertString:=func(t *testing.T, want string, got string, ffor string){

		if want != got{
			t.Errorf("want %q, got %q, for %q", want, got ,ffor)
		}
	}

	assertError:= func(t *testing.T, err error){
		if err == nil{
			t.Errorf("Expected error, didn't get one")
		}
	}
	
	t.Run("testing search for existing key", func(t *testing.T){
		dictionary := Dictionary{"test":"this is just a test"}
		definition, _:=dictionary.Search("test")
		want:="this is just a test"
		assertString(t, want, definition, "test")
})

t.Run("Seaching for non-existing key", func (t *testing.T){
	dictionary := Dictionary{"text": "this is just a test"}
	_, err := dictionary.Search("Not existant")

	assertError(t, err) 
})
}

