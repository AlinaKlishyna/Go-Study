package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 5

	res := multiply(a, b)
	fmt.Println(res)

	print("Hello")

	fmt.Println(minimumFromFour())

	fmt.Println(vote(0, 0, 1))

	fmt.Println(sumInt(1, 0))

	test(&a, &b)
}

func multiply(a, b int) int {
	return a * b
}

func print(text string) {
	fmt.Println(text)
}

func minimumFromFour() int {
	slice := make([]int, 4)
	for idx := range slice {
		fmt.Scan(&slice[idx])
	}

	min := slice[0]
	for _, elem := range slice {
		if elem < min {
			min = elem
		}
	}

	return min
}

func vote(x int, y int, z int) int {
	//print your code here
	if x == 1 && (y == 1 || z == 1) || y == 1 && (x == 1 || z == 1) || z == 1 && (x == 1 || y == 1) {
		return 1
	}
	return 0
}

func sumInt(a ...int) (int, int) {
	var sum int
	for _, elem := range a {
		sum += elem
	}
	return len(a), sum
}

func test(x1 *int, x2 *int) {
	// здесь ваш код
	save := *x1
	*x1 = *x2
	*x2 = save

	fmt.Println(*x1, *x2)
}
