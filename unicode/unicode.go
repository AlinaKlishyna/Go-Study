package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// функции ниже принимают на вход тип rune
func main() {
	// проверка символа на цифру
	fmt.Println(unicode.IsDigit('1')) // true

	// проверка символа на букву
	fmt.Println(unicode.IsLetter('a')) // true

	// проверка символа на нижний регистр
	fmt.Println(unicode.IsLower('A')) // false

	// проверка символа на верхний регистр
	fmt.Println(unicode.IsUpper('A')) // true

	// проверка символа на пробел
	// пробел это не только ' ', но и:
	//  '\t', '\n', '\v', '\f', '\r' - подробнее читайте в документации
	fmt.Println(unicode.IsSpace('\t')) // true

	// С помощью функции Is можно проверять на кастомный RangeTable:
	// например, проверка на латиницу:
	fmt.Println(unicode.Is(unicode.Latin, 'ы')) // false

	// функции преобразований
	fmt.Println(string(unicode.ToLower('F'))) // f
	fmt.Println(string(unicode.ToUpper('f'))) // F

	// len() возвращает количество байт которое занимает строка
	var en = "english"
	var ru = "русский"
	fmt.Println(len(en), len(ru))
	fmt.Println(utf8.RuneCountInString(en), utf8.RuneCountInString(ru))
}
