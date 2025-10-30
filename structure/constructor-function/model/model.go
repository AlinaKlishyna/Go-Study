package model

import "errors"

// struttura
type Player struct {
	// i campi
	Name     Name
	Birthday int
	//struttura anonima
	Location struct {
		Adress  string
		Country string
	}

	// Если вдруг добавим новое поле, то нужно будет
	// изменить конструктор
}

// структура именованная
type Name struct {
	FirstName string
	LastName  string
}

// !! функция-конструктор
// если структура изменится, то нужно будет
// изменить только эту функцию (добавить новое поле и т.д.)
func NewPerson(firstName, lastName string, birthday int, adress, country string) (*Player, error) {
	// можно добавить валидацию полей
	// например, проверить что имя не пустое
	if firstName == "" || lastName == "" {
		return nil, errors.New("first name and last name cannot be empty")
	}
	if birthday < 1900 || birthday > 2024 {
		return nil, errors.New("birthday must be between 1900 and 2024")
	}
	if adress == "" || country == "" {
		return nil, errors.New("adress and country cannot be empty")
	}

	return &Player{
		Name: Name{
			FirstName: firstName,
			LastName:  lastName,
		},
		Birthday: birthday,
		Location: struct {
			Adress  string
			Country string
		}{
			Adress:  adress,
			Country: country,
		},
	}, nil
}
