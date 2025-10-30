package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if math.Pow(float64(c), 2) == math.Pow(float64(a), 2)+math.Pow(float64(b), 2) {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
