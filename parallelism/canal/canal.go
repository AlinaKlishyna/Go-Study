package main

import (
	"fmt"
	"time"
)

func main() {
	// example()
	// example2
	// task()
	task3()
}

func removeDuplicates(input, output chan string) {
	var prev string
	for v := range input {
		if v != prev {
			output <- v
			prev = v
		}
	}
	close(output)
}

func printer(output chan string) {
	for v := range output {
		fmt.Println(v)
	}
}

func task3() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)
	go printer(outputStream)

	for _, i := range "112334456" {
		inputStream <- string(i)
	}

	close(inputStream)
}

func task2(s chan string, stroke string) {
	for i := 0; i < 5; i++ {
		s <- stroke + " "
	}
}

func task(c chan int, n int) {
	c <- n + 1
}

func example2() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"

	// periodic block/unblock of main goroutine until chanel closes
	for {
		if val, ok := <-c; !ok {
			fmt.Println(val, ok, " <-- loop broke!")
			break
		}

	}

	fmt.Println("main() stopped")
}

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
	close(c)
}

func example() {
	// Канал
	message := make(chan string) // канал для строк

	go func() {
		for i := 1; i <= 10; i++ {
			message <- fmt.Sprintf("%d", i)
			time.Sleep(time.Millisecond * 500)
		}

		close(message)
	}()

	for {
		msg, open := <-message
		if !open {
			break
		}
		fmt.Println(msg)
	}

	for msg := range message {
		fmt.Println(msg)
	}
}
