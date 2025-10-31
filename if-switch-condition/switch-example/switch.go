package main

import "fmt"

func main() {
	var i int = 1

	// Нет необходимости каждый раз писать break
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}

	// Но если мы хотим чтобы выполнение перешло к следующему коду - fallthrough
	// fallthrough - выполняется следующий блко не проверяя условие case
	switch i {
	case 0:
		fmt.Println("Zero")
		fallthrough
	case 1:
		fmt.Println("One") // Выполянется
		fallthrough
	case 2:
		fmt.Println("Two") // Выполняется
	}

	// Произвольные условия
	var z uint32
	fmt.Scan(&z)
	switch { // Не указываем переменную после switch
	case 1 <= z && z <= 9:
		fmt.Println("от 1 до 9")

	case z >= 100 && z <= 250:
		fmt.Println("от 100 до 250")

	case z >= 1000 && z <= 6000:
		fmt.Println("от 1000 до 6000")
	}

	var a int
    fmt.Scan(&a)
    
    switch {
        case a>0: 
            fmt.Println("Число положительное")
        case a<0:
            fmt.Println("Число отрицательное")
        case a==0: 
            fmt.Println("Ноль")
    }

}
