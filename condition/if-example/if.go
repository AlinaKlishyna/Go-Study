package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2
	n := 3
	limit := 20

	if v := math.Pow(float64(x), float64(n)); v < float64(limit) {
		fmt.Printf("v(%.2f) < limit(%v)\n", v, limit)
	} else if v := math.Pow(float64(x), float64(n)); v < float64(limit)+5 {
		fmt.Printf("v(%.2f) < limit(%v)\n", v, limit+5)
	} else {
		fmt.Printf("v(%.2f) > limit(%v)\n", v, limit)
	}
}
