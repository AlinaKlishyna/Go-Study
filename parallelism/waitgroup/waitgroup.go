package main

import (
	"fmt"
	"sync"
)

func main() {
	wait()
}

func wait() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1) // кол-во задач (wg++)

		go func(i int) { // anonymous function
			fmt.Println(i + 1)

			wg.Done() // таска которая была добавлена стала завершенной
		}(i) // передаем i в горутину
	}
	// ждет пока не завершится
	wg.Wait() // блокирует выполнение дальше пока все задачи не будут завершены
	fmt.Println("All tasks completed")
}
