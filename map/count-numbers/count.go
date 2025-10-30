package main

import "fmt"

func main() {

	fmt.Println(CountMaxFrequency([]int{1, 3, 2, 3, 4, 3, 2, 1})) // 3
	fmt.Println(CountMaxFrequency([]int{1, 2, 2, 3, 3}))          // 2
}

func CountMaxFrequency(numbers []int) int {
	unique := make(map[int]int)

	for _, value := range numbers {
		_, ok := unique[value]
		if !ok {
			unique[value] = 1
		} else {
			unique[value] += 1
		}
	}

	max := 0
	for k := range unique {
		if max < unique[k] {
			max = unique[k]
		}
	}
	return max
}
