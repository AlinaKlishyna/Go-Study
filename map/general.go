package main

import "fmt"

func main() {
	// где string - ключ, а int - значение
	var users map[int]int // Так писать не правильно!!
	//users[1] = 1 // panic: assignment to entry in nil map
	fmt.Println(users)

	// С использованием встроенной функции make:
	m1 := make(map[int]int)

	// Или с использованием литерала отображения:
	m2 := map[int]int{
		12: 2, // ключ 12, значение 2
		1:  5, // ключ 1, значение 5
		5:  0,
		4:  4,
	}

	fmt.Println(m1) // вывод: map[]
	fmt.Println(m2) // вывод: map[1:5 12:2]

	fmt.Println(m2[12]) // Удаляем элемент с ключом 12
	delete(m2, 1)       // Удаляем элемент с ключом 12

	// Важно, что операции обращения и удаления в отображении безопасны, даже если ключ не существует
	// Если мы попытаемся обратиться к ключу, которого нет в отображении, то Go вернет нулевое значение
	// соответствующего типа
	m := make(map[int]int)
	fmt.Println(m[12]) // 0 (по умолчанию int = 0)
	delete(m, 12)
	fmt.Println(m) // map[]

	// Чтобы избежать нужно использовать проверку наличия ключа
	if value, inMap := m2[4]; inMap {
		fmt.Println("-->", value) // Условие не выполнится
	}

	if value, inMap := m2[12]; inMap {
		fmt.Println("-->", value) // Условие не выполнится
	}

	// перебирать все элементы
	for key, value := range m2 {
		fmt.Println(key, " - ", value)
	}

	// len() для определения количества пар "ключ-значение"
	fmt.Println("lenght m2: ", len(m2))

	// Основные операции
	// Вставка
	m2[5] = 20

	// Удаление
	delete(m2, 5)

	dispalyCreateMap() // вывод разнится (мапа не упорядоченная)
	search()
}

// мапа в Go unordered, то есть не упорядоченная
func dispalyCreateMap() {
	m := map[int]bool{}
	for i := 0; i < 5; i++ {
		m[i] = i%2 == 0
	}

	for k, v := range m {
		fmt.Printf("key: %d - value: %t\n", k, v)
	}
}

// Поиск в таблице Go
func search() {
	mapa := map[int]int{0: 0, 1: 10}
	value, ok := mapa[2]
	if !ok {
		mapa[2] = 20
	}
	fmt.Println(mapa, mapa[0], mapa[1], value)
}
