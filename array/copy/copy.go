package main

import "fmt"

func main() {
	first := []int{1, 4, 7, 9}
	second := make([]int, len(first), cap(first))
	result := copy(second, first)

	fmt.Printf("first = %v\n", first)
	fmt.Printf("second = %v\n", second)
	fmt.Println(result) // Возвращает число скопированных элементов

	a := [3]int{1, 2, 3} // массив
	b := []int{1, 2, 3}  // срез

	fnA(a)
	fnB(b)

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 15 3]
}

// [3]int - массив
// когда передаем массив в функцию, то на самом деле создается копия
// и меняется только внутри данного блока
func fnA(a [3]int) { // изменяет значение по индексу 1 ВНУТРИ функции
	a[1] = 15
}

// []int - не массив, а структура с указателем на реальный массив
func fnB(a []int) { // ИЗМЕНЯЕТ САМ МАССИВ
	a[1] = 15
}
