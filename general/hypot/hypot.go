package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	hypot := math.Hypot(float64(a), float64(b)) // Гипотенуза
	fmt.Println(hypot)
}
