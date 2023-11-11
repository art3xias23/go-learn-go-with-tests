package arrays

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("Testing a sum of array with 5 elements", func(t *testing.T) {
		numbers := []int{4, 5, 6, 7}

		got := Sum(numbers)
		expected := 22

		if expected != got {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("Testing a sum of array with any size of elements", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if expected != got {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
