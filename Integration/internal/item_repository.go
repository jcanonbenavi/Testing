package internal

import "errors"

var (
	// ErrRepositoryNotFound is returned when an item is not found in the database
	ErrRepositoryNotFound = errors.New("repository: item not found")
	// ErrRepositoryInternal is returned when an internal error occurs
	ErrRepositoryInternal = errors.New("repository: internal error")
)

type ItemRepository interface {
	FindById(id int) (i Item, err error)
}
