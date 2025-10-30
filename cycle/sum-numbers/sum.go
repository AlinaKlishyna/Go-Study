package main

import "fmt"

func main() {
	fmt.Println(sumAll(1, 5))
}

func sumAll(a, b int) int {
	var sum int
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}
