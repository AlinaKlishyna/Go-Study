package main

import "fmt"

func main() {
	fmt.Println("Hello! This is my first start by Alina!") // Выводит строку с переносом
	fmt.Print("Bye!")                                      // Выводит строку без переноса

	// Использование конкатенации строк через " " + " "
	fmt.Println(len("Bye!" + "й")) // Выводит количество байт. Русские буквы занимают 2 байта

	// По индексу ищем байт строки>
	fmt.Println("Alina"[0])
}
