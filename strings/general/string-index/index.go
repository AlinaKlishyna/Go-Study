package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindIndexFirstInput("awesome", "es"))
}

func FindIndexFirstInput(x, s string) int {
	return strings.Index(x, s)
}
