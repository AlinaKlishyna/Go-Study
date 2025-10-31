package main

import "fmt"

type Hero struct {
	On    bool
	Ammo  int
	Power int
}

func (t *Hero) Shoot() bool {
	if t.On == false {
		return false
	} else if t.Ammo > 0 {
		t.Ammo--
		return true
	}

	return false
}

func (t *Hero) RideBike() bool {
	if t.On == false {
		return false
	} else if t.Power > 0 {
		t.Power--
		return true
	}
	return false
}

func main() {
	person := Hero{On: true, Ammo: 250, Power: 15}
	fmt.Println(person)
	person.Shoot()
	fmt.Println(person)
	person.RideBike()
	fmt.Println(person)
}
