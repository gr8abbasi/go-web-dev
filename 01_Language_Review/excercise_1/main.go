package main

import "math"
import "fmt"

// Shape Interface
type Shape interface {
	area() float64
}

// Square Struct
type Square struct {
	height float64
	width  float64
}

func (s Square) area() float64 {
	return s.height * s.width
}

// Circle Struct
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	square := Square{3, 6}
	circle := Circle{4}

	info(square)
	info(circle)
}

func info(s Shape) {
	fmt.Printf("Area for %T is %f :\n", s, s.area())
}
