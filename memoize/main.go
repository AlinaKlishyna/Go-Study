package main

import (
	"fmt"
	"time"
)

func slowOperation(i int) int {
	time.Sleep(1 * time.Second)
	return i + 1
}

func main() {
	memoziedOperation := memoize(slowOperation)

	fmt.Println(memoziedOperation(1))
	fmt.Println(memoziedOperation(1))
	fmt.Println(memoziedOperation(1))
}

func memoize(
	expensiveCalculation func(i int) int,
) func(i int) int {
	cache := make(map[int]int)

	return func(i int) int {
		cachedResult, ok := cache[i]
		if ok {
			return cachedResult
		}

		result := expensiveCalculation(i)
		cache[i] = result

		return result
	}
}
