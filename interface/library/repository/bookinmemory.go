package repository

import (
	"fmt"
	"interface/library/model"
)

var (
	ErrNotFoundID = fmt.Errorf("Book already exists")
)

type BookInMemoryRepository struct {
	books map[int]model.Book
}

func NewBookInMemoryRepository() *BookInMemoryRepository {
	return &BookInMemoryRepository{
		books: make(map[int]model.Book),
	}
}

func (r BookInMemoryRepository) ByID(ID int) (model.Book, error) {
	book, exists := r.books[ID]
	if !exists {
		// %w - error
		return model.Book{}, fmt.Errorf("%w: ID %d", ErrNotFoundID, ID)
	}

	return book, nil
}

func (r BookInMemoryRepository) Add(b model.Book) error {
	if _, exists := r.books[b.ID]; exists {
		return ErrErrNotFoundID
	}
	r.books[b.ID] = b
	return nil
}

func (r BookInMemoryRepository) Delete(ID int) error {
	if _, exists := r.books[b.ID]; exists {

	}
}

// DA FARE
