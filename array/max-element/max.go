package main

import "fmt"

func main() {
	array := [5]int{}

	for idx := range array {
		fmt.Scanln(&array[idx])
	}

	max := array[0]

	for _, elem := range array {
		if max < elem {
			max = elem
		}
	}

	fmt.Println(max)
}
