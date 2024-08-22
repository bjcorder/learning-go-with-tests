package perimeter

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	perimeter := 2 * (rectangle.width + rectangle.height)
	return perimeter
}

func Area(rectangle Rectangle) float64 {
	area := rectangle.width * rectangle.height
	return area
}
