package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Printf("%s speaking..\n", a.Name)
}

// embedding
// встраивание позволяет создавать новые структуры,
// используя существующие, что позволяет повторно использовать код и упрощает создание сложных типов
// Мы объединяем несколько структур в более сложную структуру, комбинируя функциональность различных типов
type Dog struct {
	// Собака ЭТО животное
	Animal // это встраивание структуры Animal. Теперь Dog имеет все поля и методы Animal.
	Breed  string
}

func (d Dog) Bark() {
	fmt.Println("Woof!")
}

func main() {
	dog := Dog{
		Animal: Animal{
			Name: "Sharik",
		},
		Breed: "Retriver",
	}

	fmt.Printf("%+v\n", dog)

	dog.Speak()                             // Вызов метода Speak из встроенной структуры Animal
	dog.Bark()                              // Вызов метода Bark из структуры Dog
	fmt.Println("Dog's name is:", dog.Name) // Доступ к полю Name из встроенной структуры Animal

	dog.Name = "Bobik" // Изменение поля Name из встроенной структуры Animal
	fmt.Println("Dog's new name is:", dog.Name)
	fmt.Printf("%+v\n", dog)

}
