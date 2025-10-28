package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	d := [3]int{1: 12} //[количество индексов] int {индекс: число}

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12 0]

	// Массивы так же можно сравнивать между собой
	fmt.Println(a == b)
	fmt.Println(c == d)

	// Достать индекс масива - [0]
	var numbers [5]int = [5]int{2, 5, 6, 7, 5}
	fmt.Println(numbers[0])
	fmt.Println(numbers[2])

	// Вывод массива
	var num [5]int = [5]int{3, 2, 5, 6, 8} // [3 2 5 6 8]
	fmt.Println(num)

	for i := 0; i < len(num); i++ {
		fmt.Println(num[i]) // len() - возвращает длинну массива
		// 3
		// 2
		// 5
		// 6
		// 8
	}

	// Вывод элементов массива
	for idx, elem := range num {
		fmt.Printf("Элемент c индексом %d: %d\n", idx, elem)
	}

	// Вывод элементов через индекс
	for idx := range num {
		fmt.Println(num[idx])
	}

	// Вывод элементов
	for _, elem := range num {
		fmt.Println(elem)
	}

	// Чтобы изменить элементы массив, то только через индекс
	for _, elem := range num {
		elem = 100
		fmt.Println(elem)
		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(num) // массив НЕ ИЗМЕНИЛСЯ

	for idx, elem := range num {
		num[idx] = 100
		// 100
		// 100
		// 100
		// 100
		// 100
		fmt.Println(elem) // теперь массив ИЗМЕНИЛСЯ
	}
}
