package main

import (
	"fmt"
)

// встраивание (embedding) (A) позволяет обращаться к полям и методам вложенной структуры напрямую
// как если бы они были частью внешней структуры

// В случае композиции (A A) доступ к полям осуществляется через поле, которое содержит вложенную структуру
type A struct {
	Property string
}

// composition -- Композиция позволяет создавать новые структуры,
// используя существующие структуры в качестве полей
type Composite struct {
	A A // когда это всего лишь часть
	// Engine isn't Car
}

type Embedded struct {
	A // когда А неотьемлемая часть структуры  --> Emedded это A
	// Car is Autobus
}

func main() {
	a := A{Property: "Hello"}

	composite := Composite{A: a}

	embed := Embedded{
		A: a,
	}

	fmt.Printf("%+v\n", composite)
	fmt.Printf("%+v\n", embed)

	fmt.Println(composite.A.Property)
	fmt.Println(embed.A.Property)

	// fmt.Println(composite.Property) - !! Нельзя взять напрямую в композиции
	fmt.Println(embed.Property)
}
