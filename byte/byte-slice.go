package main

import (
	"fmt"
)

func main() {
	// Берем строку и конвертируем в массив байтов
	byteSlice := []byte("Это байтовый срез")

	fmt.Println(byteSlice)

	for idx := range byteSlice {
		if idx%2 == 0 {
			byteSlice[idx] += 1
		} else {
			byteSlice[idx] -= 1
		}

	}

	fmt.Println(string(byteSlice))
}
