package main

import "fmt"

type Dog struct{}

func (Dog) Speak() string {
	return "Woof!"
}

func (Dog) Growl() string {
	return "RRRrrRRrrRRRR!!!"
}

type Cat struct{}

func (Cat) Speak() string {
	return "Meow-meow.."
}

type Bird struct{}

func (Bird) Speak() string {
	return "Cirik-cirik"
}

func main() {
	cat := Cat{}
	dog := Dog{}
	bird := Bird{}

	MakeSound(cat)
	MakeSound(dog)
	MakeSound(bird)
}

type Speaker interface {
	Speak() string
}

func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}
