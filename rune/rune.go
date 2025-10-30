package main

import "fmt"

func main() {
	// Преобразуем строку в срез рун
	rs := []rune("Это срез рун")

	fmt.Println(rs)

	for idx, elem := range rs {
		if elem == ' ' {
			rs[idx] = '.'
		}
	}

	fmt.Println(string(rs))
}
