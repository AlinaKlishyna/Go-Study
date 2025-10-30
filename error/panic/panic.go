package main

import (
	"fmt"
)

// если есть panic и defer, то сначала выполняется defer, а потом panic
func main() {
	// Оператор panic позволяет сгенерировать ошибку и выйти из программы
	fmt.Println(divide(10, 0))
	// После этого выполнение программы будет остановлено, и все последующие операции не будут выполнены
	fmt.Println("После panic")

	// panic: Division by Zero!
}

func divide(a, b int) (int, error) {
	if b == 0 {
		panic("Division by Zero!")
	}
	return a / b, nil
}
