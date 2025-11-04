package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// wait()
	writeWithoutMutex()
	writeWithtMutex()
}

func wait() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1) // кол-во задач (wg++)

		go func(i int) { // anonymous function
			fmt.Println(i + 1)

			defer wg.Done() // отложенный вызов, который выполнится в конце функции

			fmt.Println("i: ", i+10)
		}(i) // передаем i в горутину
	}
	// ждет пока не завершится
	wg.Wait() // блокирует выполнение дальше пока все задачи не будут завершены
	fmt.Println("All tasks completed")
}

func writeWithoutMutex() {
	start := time.Now()
	var wg sync.WaitGroup
	var counter int

	wg.Add(1000)

	for i := 1; i <= 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			// Data Race - обращение к одной и той же переменной из нескольких горутин,
			// что приводит к непредсказуемым результатам
			counter++ // counter = counter + 1 || counter + 1  = 556 || counter + 1 = 556
		}()
	}

	wg.Wait()
	fmt.Println("Counter without mutex:", counter)
	fmt.Println("Time without mutex:", time.Now().Sub(start).Seconds())
}

func writeWithtMutex() {
	start := time.Now()
	var wg sync.WaitGroup
	var counter int
	var mu sync.Mutex

	wg.Add(1000)

	for i := 1; i <= 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			// Data Race - обращение к одной и той же переменной из нескольких горутин,
			// что приводит к непредсказуемым результатам

			// использование мьютекса для синхронизации доступа к ресурсу
			mu.Lock()   // блокируем доступ к ресурсу
			counter++   // counter + 1  = 557 || counter + 1 = 558
			mu.Unlock() // разблокируем доступ к ресурсу
		}()
	}

	wg.Wait()
	fmt.Println("Counter without mutex:", counter)
	fmt.Println("Time without mutex:", time.Now().Sub(start).Seconds())
}
