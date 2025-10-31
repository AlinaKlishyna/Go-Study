package main

import "fmt"

func main() {
	var text string
	println("Введите текст:")
	fmt.Scan(&text)

	fmt.Println(strokeStars(text))
}

func strokeStars(text string) string {
	var update string
	for idx, letter := range text {
		update += string(letter)

		if idx < len(text)-1 {
			update += "*"
		}
	}
	return update
}
