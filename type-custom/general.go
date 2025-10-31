package main

import "fmt"

// enums
type Direction int

const (
	North Direction = iota + 1
	East
	South
	West
)

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	default:
		panic("Unknown direction")
	}
}

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	}
	return ""
}

func main() {
	fmt.Println(South)

	action(South)
	actionToStroke(South)

}

func action(d Direction) {
	fmt.Println("Action", d)
}

func actionToStroke(d fmt.Stringer) {
	fmt.Println("Action", d)
}
