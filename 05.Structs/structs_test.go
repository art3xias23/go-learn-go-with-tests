package structs

import "testing"

func TestShapes(t *testing.T) {
	t.Run("should get the area of a rectangle", func(t *testing.T) {
		checkErrors(t, Rectangle{4,5}, 18.0)

	})

	t.Run("should get the area of a rectangle", func(t *testing.T) {
		checkErrors(t, Circle{10}, 314.1592653589793)
	})
}

func checkErrors(t *testing.T, shape Shape, want float64){
	got := shape.Area()
		if want != got {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
}
