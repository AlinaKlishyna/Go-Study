package main

import "fmt"

func main() {
	var a float64 = 100.123456789
	result := fmt.Sprintf("%.2f", a)
	// %q вывод в одинарных кавычках " "
	fmt.Printf("%q", result) // вывод: "100.12"
	// result будет типа string
}
