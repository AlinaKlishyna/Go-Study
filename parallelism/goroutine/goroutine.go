package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// сколько физических ядер процессора можно использовать одновременно
	// По умолчанию Go уже сам использует все доступные ядра, но раньше (в старых версиях Go) нужно было указывать вручную
	runtime.GOMAXPROCS(4)

	// go — это ключевое слово, запускающее новую горутину
	go func() { // анонимная функция, то есть функция без имени
		fmt.Println("I'm am anonim function!")
	}() // () в конце — выполняют (вызывают) анонимную функцию сразу же
	time.Sleep(1 * time.Second)

	fmt.Println(runtime.NumCPU()) // максимальное кол-во goroutine

	go num() // Запускает функцию num() в отдельной горутине — то есть параллельно

	runtime.Gosched() // предлагает Go дать другим поработать, переключатель
	time.Sleep(time.Second)
	fmt.Println("exit..")
}

func num() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}
