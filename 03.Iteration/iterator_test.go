package iterator

import "testing"

func TestIterate(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
