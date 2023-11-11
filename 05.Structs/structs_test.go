package structs

import "testing"

func TestShapes(t *testing.T) {
	t.Run("should get the area of a rectangle", func(t *testing.T) {
		got := Rectangle{4, 5}.Area()
		want := 18.0

		if want != got {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	})

	t.Run("should get the area of a rectangle", func(t *testing.T) {
		got := Circle{10}.Area()
		want := 314.1592653589793
		if want != got {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	})
}
