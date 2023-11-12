package main

import "testing"

func TestSearch(t *testing.T){

	assertError:=func(t *testing.T, want string, got string, ffor string){

		if want != got{
			t.Errorf("want %q, got %q, for %q", want, got ,ffor)
		}
	}
	
	t.Run("testing search for key", func(t *testing.T){
		dictionary := map[string]string{"test":"this is just a test"}
		got:=Search(dictionary, "test")
		want:="this is just a test"
		assertError(t, want, got, "test")
})
}

