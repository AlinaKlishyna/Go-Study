package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(WithoutDuplicate("abacabad"))  // abcd
	fmt.Println(WithoutMoreSymbols("zaabcbd")) // zcd
}

func WithoutDuplicate(text string) string {
	var update string

	for idx := range text {
		if !strings.Contains(update, string(text[idx])) {
			update += string(text[idx])
		}
	}
	return update
}

func WithoutMoreSymbols(text string) string {
	var update string

	for _, elem := range text {
		if strings.Count(text, string(elem)) == 1 {
			update += string(elem)
		}
	}
	return update
}
