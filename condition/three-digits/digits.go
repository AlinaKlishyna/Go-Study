package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var primo = a % 10
	a /= 10
	var secondo = a % 10
	a /= 10
	var terzo = a % 10

	if primo != secondo && primo != terzo && secondo != terzo {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
