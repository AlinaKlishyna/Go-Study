package main

import "errors"

// DA FARE

type User struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

func (u User) SecretIdentity() string {
	return ""
}

func (u User) Age() int {
	return 0
}

func (u User) AddFavoriteLanguage(language string) error {
	return errors.New("not realized")
}

func (u User) RemoveFavoriteLanguage(language string) error {
	return errors.New("not realized")
}

func (u User) IsProgrammingLanguageFavorite(language string) bool {
	return false
}

func (u User) RandomFavoriteLanguage() (string, error) {
	return "", errors.New("not realized")
}

func (u User) GenerateProfile() string {
	return ""
}

func (u User) UpdateName(firstName, lastName string) error {
	return errors.New("not realized")
}
