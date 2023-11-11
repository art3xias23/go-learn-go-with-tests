package arrays

import (
	"reflect"
	"testing"
)

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

func TestSumAllTails(t *testing.T) {
	t.Run("Run populated slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})

	t.Run("Safely run empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3, 4})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
