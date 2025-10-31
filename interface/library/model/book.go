package model

import (
	"errors"
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author string
}

var (
	ErrIncorrectID = errors.New("ID must be greater 0")
	ErrEmptyTitle  = errors.New("Title cannot be empty")
	ErrEmptyAuthor = errors.New("Author cannot be empty")
)

func NewBook(ID int, title string, author string) (Book, error) {
	if ID <= 0 {
		return Book{}, ErrIncorrectID
	}

	if title == "" {
		return Book{}, ErrEmptyTitle
	}

	if author == "" {
		return Book{}, ErrEmptyAuthor
	}

	return Book{
		ID:     ID,
		Title:  title,
		Author: author,
	}, nil
}

func (b Book) String() string {
	return fmt.Sprintf("ID: %d,\nName: %s,\nAuthor: %s", b.ID, b.Title, b.Author)
}
