package structs

type Shape interface {Area() float64}
type Circle struct{ Radius float64 }

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.141592653589793
}

func (r Rectangle) Area() float64 {
	return r.Width*2 + r.Height*2
}
