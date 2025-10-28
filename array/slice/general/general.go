package main

import "fmt"

func main() {
	numbers := []int{10, 11, 12, 13, 14}
	numbers = append(numbers, 15, 16)
	fmt.Println(numbers)

	// Создание слайса через make
	slice := make([]int, 3, 5)
	fmt.Println(slice) // [0 0 0]
	// len() - длинна массива cap() - емкость
	fmt.Println(len(slice), cap(slice)) // 3 5

}
