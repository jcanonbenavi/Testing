package internal

import "errors"

var (
	// ErrServiceNotFound is returned when an item is not found in the database
	ErrServiceNotFound = errors.New("service: item not found")
	// ErrServiceInternal is returned when an internal error occurs
	ErrServiceInternal = errors.New("service: internal error")
)

type ItemService interface {
	FindById(id int) (i Item, err error)
}
