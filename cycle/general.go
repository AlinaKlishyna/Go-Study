package main

import "fmt"

func main() {
	sum := 0

	// Первый вариант цикла
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	// Второй вариант цикла - вынести переменную вне блока
	var i = 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	// Можно оставить только условие цикла
	// Но увеличить внутри
	var j = 1
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Считываем числа до тех пор, пока не введен 0
	// но в таком виде ошибки ввода не обрабатываются
	var num int
	for fmt.Scan(&num); num != 0; fmt.Scan(&num) {
		fmt.Println(num)
	}
}
