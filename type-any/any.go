package main

import "fmt"

// использование типа any требует явной проверки типа (type assertion) для работы с данными,
// что может привести к ошибкам времени выполнения и усложняет код
func main() {
	// type any = interface{}
	var a any // Пустой интерфейс может хранить значение любого типа
	Print(a)  // <nil>
	PrintElements([]any{43, "Hi", '±', nil, 3.14})
}

func Print(a any) {
	// Println - func fmt.Println(a ...any) - принимает переменное количество аргументов любого типа
	fmt.Println(a)
}

func PrintElements(a []any) {
	for i := 0; i < len(a); i++ {
		if _, ok := a[i].(int); ok {
			fmt.Printf("Элемент %d является числом: %d\n", i, a[i])
		} else if _, ok := a[i].(string); ok {
			fmt.Printf("Элемент %d является строкой: %s\n", i, a[i])
		} else if _, ok := a[i].(rune); ok {
			fmt.Printf("Элемент %d является руной: %c\n", i, a[i])
		} else {
			fmt.Printf("Элемент %d имеет тип %T и значение %v\n", i, a[i], a[i])
		}
	}
}
