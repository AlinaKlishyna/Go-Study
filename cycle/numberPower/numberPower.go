package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(math.Pow(float64(i), 2))
	}
}
