package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// [n]int - НЕ МОЖЕМ ТАК СДЕЛАТЬ ПОТОМУ ЧТО КОМПИЛЯТОР НЕ ЗНАЕТ ДО ЗАПУСКА КАКОЕ ЧИСЛО НАПИШЕМ ЧЕЛОВЕК
	workArray := make([]int, n, 10)
	// range - позволяет пройтись по елементам
	// на каждой итерации idx - получает индекс массива
	// но значение елемента мы не берем (elem не указано)
	for idx := range workArray {
		fmt.Scan(&workArray[idx])
	}

	fmt.Println(workArray[3])

}
