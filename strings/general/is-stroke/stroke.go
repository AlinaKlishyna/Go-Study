package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	stroke, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	isStroke(stroke)
}

func isStroke(text string) {
	text = strings.TrimSuffix(text, "\n")

	symbols := []rune(text)
	output := "Wrong"
	if !unicode.IsLower(symbols[0]) {
		if symbols[len(symbols)-1] == '.' {
			output = "Right"
		}
	}
	fmt.Println(output)
}
