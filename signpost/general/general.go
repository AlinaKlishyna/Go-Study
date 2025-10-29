package main

import "fmt"

func main() {
	a := 200
	// & - передает адресс переменной
	b := &a //b - здесь является указателем
	*b++
	fmt.Println(a)
}
