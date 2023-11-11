package arrays
import "testing"

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
