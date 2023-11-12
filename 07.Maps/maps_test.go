package main

import "testing"

func TestMap(t *testing.T){

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

t.Run("should find an item after inserting it" ,func(t *testing.T){
		var d = Dictionary{}
		d.Add("test", "this is test")
		def, _ := d.Search("test")
		want:="this is test"
		assertString(t, want, def, "test")
	})
}

