package main

import "fmt"

func main() {
	var workArray [10]int

	for idx := range workArray {
		fmt.Scan(&workArray[idx]) // записываем значения с консоли напрямую в массив
	}

	fmt.Println(workArray)
}
