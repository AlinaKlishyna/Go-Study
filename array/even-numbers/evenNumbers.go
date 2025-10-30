package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	array := make([]int, n)

	for idx := range array {
		fmt.Scan(&array[idx])
	}

	for idx, elem := range array {
		if idx%2 == 0 {
			fmt.Print(elem, " ")
		}
	}
}
