package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name   string
	Grades []int
}

// Значения не меняются!
//
//	func (s Student) ChangeName(name string) {
//	   s.Name = name
//	}

// ЕСЛИ ДАННЫЕ МЕНЯЕМ ТО ПИЩЕМ *  --> (s *Student)
func (s *Student) ChangeName(name string) {
	s.Name = name
}

// ЕСЛИ ДАННЫЕ НЕ МЕНЯЕМ НЕ ПИШЕМ  --> (s Student)
func (s Student) AverageGrade() float64 {
	var average float64
	if len(s.Grades) == 0 {
		return 0
	}
	for _, v := range s.Grades {
		average += float64(v)
	}
	average /= float64(len(s.Grades))
	return math.Round(average*10) / 10 // округляем до 1 знака после запятой
}

func (s Student) Info() string {
	return fmt.Sprintf("Студент %s, средняя оценка: %.1f.", s.Name, s.AverageGrade())
}

func main() {
	student := Student{
		Name:   "Alina",
		Grades: []int{3, 5, 4, 2, 2},
	}

	fmt.Println(student.Info()) // Студент Alina, средняя оценка: 3.2.

	student.ChangeName("Pavel")

	fmt.Println(student.Info()) // Студент Pavel, средняя оценка: 3.2.

	str := "hello"

	fmt.Println([]byte(str))
	fmt.Println([]rune(str))
}
