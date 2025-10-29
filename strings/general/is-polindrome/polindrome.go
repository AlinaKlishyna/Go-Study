package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPolindrome("топот"))
}

func isPolindrome(text string) string {
	symbols := []rune(text)
	var reverse string
	for i := len(symbols) - 1; i >= 0; i-- {
		reverse += string(symbols[i])
	}
	if text == reverse {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
	return reverse
}
