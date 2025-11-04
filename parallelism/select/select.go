package main

import "fmt"

func main() {
	baseSelect()
}

func baseSelect() {
	bufferedChan := make(chan string, 1) // буферизированный канал с размером буфера 1
	bufferedChan <- "first"              // записываем в канал первое значение

	select {
	case stroke := <-bufferedChan: // успешно читает из канала
		fmt.Println("read", stroke)
	case bufferedChan <- "second": // блокируется, т.к. канал заполнен
		fmt.Println("write", <-bufferedChan, <-bufferedChan)
	}

}
