package iterator

import "testing"

func TestIterate(t *testing.T) {
	t.Run(
		"Tetsing 4 repeating chars", func(t *testing.T) {
			got := Repeat("a", 4)
			expected := "aaaa"

			if got != expected {
				t.Errorf("expected %q, got %q", expected, got)
			}
		})

	t.Run("Testing with 6 repeating chars", func(t *testing.T) {
		got := Repeat("b", 6)
		expected := "bbbbbb"

		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})
}
