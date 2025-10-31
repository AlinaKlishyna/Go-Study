package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width * r.height)
}

func (r Rectangle) PrintInfo() {
	fmt.Printf("Shape %s,\nwidth: %.2f,\nheight: %.2f,\narea: %.2f,\nperimeter: %.2f\n\n", "Rectangle", r.width, r.height, r.Area(), r.Perimeter())
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) PrintInfo() {
	fmt.Printf("Shape %s,\nradius: %.2f,\narea: %.2f\nperimeter: %.2f\n\n", "Circle", c.radius, c.Area(), c.Perimeter())
}

func main() {
	rect := Rectangle{
		width:  10,
		height: 5,
	}
	cirle := Circle{radius: 10}

	// rect.PrintInfo()
	// cirle.PrintInfo()
	PrinterInfoShape(rect)
	PrinterInfoShape(cirle)
}

type Shaper interface {
	Area() float64
	Perimeter() float64
}

type Printer interface {
	PrintInfo()
}

func PrinterInfoShape(p Printer) {
	p.PrintInfo()
}
