package main

import (
	"fmt"
	"math"
)

// type -  создает новый тип данных
// struct -  структура, которая объединяет в себе несколько полей
type Circle struct {
	x float64
	y float64
	r float64

	// x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := r.x2 - r.x1
	w := r.y2 - r.y1
	return l * w
}

func main() {
	rect := Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}
	circ := Circle{x: 0, y: 2, r: 5}
	fmt.Println("Rectangle area: ", rect.area())
	fmt.Println("Circle area: ", circ.area())
}
