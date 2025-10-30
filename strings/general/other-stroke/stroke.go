package main

import "fmt"

func main() {
	var text string = "ihgewlqlkot"
	fmt.Println(DeleteOddIndex(text))
}

func DeleteOddIndex(text string) string {
	var textEvenIndex string
	for idx := range text {
		if idx%2 != 0 {
			textEvenIndex += string(text[idx])
		}
	}
	return textEvenIndex
}
