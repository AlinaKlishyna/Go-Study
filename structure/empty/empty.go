package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := struct{}{} // empty struttura

	fmt.Println(s)
	fmt.Println(unsafe.Sizeof(s)) // size in memoria 0

	numbers := []int{3, 5, 2, 5, 43}
	fmt.Println(notUnique(numbers))
	fmt.Println(notUniqueOptimized(numbers))

}
func notUniqueOptimized(s []int) bool {
	m := make(map[int]struct{}, len(s))

	for i := 0; i < len(s); i++ {
		value := s[i]
		if _, ok := m[value]; ok {
			return true
		}
		m[value] = struct{}{}
	}
	return false
}

func notUnique(s []int) bool {
	// Oˆ2 - tanto
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s)-1; j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}

	return false
}
