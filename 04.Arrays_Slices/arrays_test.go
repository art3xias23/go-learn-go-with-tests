package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{4, 5, 6, 7}

	got := Sum(numbers)
	expected := 22

	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
