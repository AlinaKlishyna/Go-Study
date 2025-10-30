package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	array := make([]int, n)

	for idx := range array {
		fmt.Scan(&array[idx])
	}

	var count int
	for _, elem := range array {
		if elem >= 0 {
			count++
		}
	}

	fmt.Println(count)

}
