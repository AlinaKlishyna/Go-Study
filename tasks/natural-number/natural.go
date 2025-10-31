package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	var count int
	for _, elem := range array {
		if elem == 0 {
			count++
		}
	}
	fmt.Println(count)
}
